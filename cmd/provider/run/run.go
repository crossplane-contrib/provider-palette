package run

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"gopkg.in/alecthomas/kingpin.v2"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	clusterapis "github.com/crossplane-contrib/provider-palette/apis/cluster"
	namespacedapis "github.com/crossplane-contrib/provider-palette/apis/namespaced"
	"github.com/crossplane-contrib/provider-palette/config"
	"github.com/crossplane-contrib/provider-palette/internal/clients"
	clustercontroller "github.com/crossplane-contrib/provider-palette/internal/controller/cluster"
	// namespacedcontroller "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced"
	"github.com/crossplane-contrib/provider-palette/internal/features"
	"github.com/crossplane-contrib/provider-palette/internal/utils"
)

const (
	tlsServerCertDir = "/var/run/secrets/tls"
	tlsServerCertDirEnvVar = "TLS_SERVER_CERTS_DIR"
	webhookTLSCertDirEnvVar = "WEBHOOK_TLS_CERT_DIR"
	certsDirEnvVar = "TLS_SERVER_CERTS_DIR"
)

var cancel context.CancelFunc

func Run() {
	var (
		app                     = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for Palette").DefaultEnvars()
		debug                   = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncInterval            = app.Flag("sync", "Sync interval controls how often all resources will be double checked for drift.").Short('s').Default("1h").Duration()
		pollInterval            = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		// pollStateMetricInterval = app.Flag("poll-state-metric", "State metric recording interval").Default("5s").Duration()
		leaderElection          = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		terraformVersion        = app.Flag("terraform-version", "Terraform version.").Required().Envar("TERRAFORM_VERSION").String()
		maxReconcileRate        = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may checked for drift from the desired state.").Default("10").Int()
		// webhookPort             = app.Flag("webhook-port", "The port the webhook listens on").Default("9443").Envar("WEBHOOK_PORT").Int()
		// metricsBindAddress      = app.Flag("metrics-bind-address", "The address the metrics server listens on").Default(":8080").Envar("METRICS_BIND_ADDRESS").String()
		// changelogsSocketPath    = app.Flag("changelogs-socket-path", "Path for changelogs socket (if enabled)").Default("/var/run/changelogs/changelogs.sock").Envar("CHANGELOGS_SOCKET_PATH").String()

		providerSource  = app.Flag("terraform-provider-source", "Terraform provider source.").Required().Envar("TERRAFORM_PROVIDER_SOURCE").String()
		providerVersion = app.Flag("terraform-provider-version", "Terraform provider version.").Required().Envar("TERRAFORM_PROVIDER_VERSION").String()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()
		// enableChangeLogs         = app.Flag("enable-changelogs", "Enable support for capturing change logs during reconciliation.").Default("false").Envar("ENABLE_CHANGE_LOGS").Bool()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))
	log.Default().SetOutput(io.Discard)
	ctrl.SetLogger(zap.New(zap.WriteTo(io.Discard)))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("provider-palette"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	// currently, we configure the jitter to be the 5% of the poll interval
	pollJitter := time.Duration(float64(*pollInterval) * 0.05)
	log.Debug("Starting", "sync-interval", syncInterval.String(),
		"poll-interval", pollInterval.String(), "poll-jitter", pollJitter, "max-reconcile-rate", *maxReconcileRate)

	// If TF_PID and TF_ADDR are set, configure the TF_REATTACH_PROVIDERS environment variable
	err := utils.ConfigureTFReattachProvider()
	kingpin.FatalIfError(err, "Cannot configure TF_REATTACH_PROVIDERS")

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:             *leaderElection,
		LeaderElectionID:           "crossplane-leader-election-provider-palette",
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")

	kingpin.FatalIfError(clusterapis.AddToScheme(mgr.GetScheme()), "Cannot add Palette cluster APIs to scheme")
	kingpin.FatalIfError(namespacedapis.AddToScheme(mgr.GetScheme()), "Cannot add Palette namespaced APIs to scheme")
	kingpin.FatalIfError(apiextensionsv1.AddToScheme(mgr.GetScheme()), "Cannot add api-extensions APIs to scheme")


	clusterOpts := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  log,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            1 * time.Minute,
			MaxConcurrentReconciles: 1,
			Features:                &feature.Flags{},
		},
		Provider: config.GetProvider(),
		// use the following WorkspaceStoreOption to enable the shared gRPC mode
		// terraform.WithProviderRunner(terraform.NewSharedProvider(log, os.Getenv("TERRAFORM_NATIVE_PROVIDER_PATH"), terraform.WithNativeProviderArgs("-debuggable")))
		WorkspaceStore: terraform.NewWorkspaceStore(log),
		SetupFn:        clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion),
	}

	// TODO: add namespaced controller PA1 
	// namespacedOpts := tjcontroller.Options{
	// 	Options: xpcontroller.Options{
	// 		Logger:                  log,
	// 		GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
	// 		PollInterval:            1 * time.Minute,
	// 		MaxConcurrentReconciles: 1,
	// 		Features:                &feature.Flags{},
	// 	},
	// }

	if *enableManagementPolicies {
		clusterOpts.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}


	kingpin.FatalIfError(clustercontroller.Setup(mgr, clusterOpts), "Cannot setup Palette controllers")
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}

func Cancel() {
	cancel()
}
