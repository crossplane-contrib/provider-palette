package integration

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	clusterapis "github.com/crossplane-contrib/provider-palette/apis/cluster"
	clusterv1beta1 "github.com/crossplane-contrib/provider-palette/apis/cluster/v1beta1"
	"github.com/crossplane-contrib/provider-palette/cmd/provider/run"
	"github.com/crossplane-contrib/provider-palette/internal/utils"
	"github.com/crossplane-contrib/provider-palette/tests/integration/routes"
)

const (
	timeout  = time.Second * 300
	interval = time.Millisecond * 250
)

var (
	kclient client.Client
	cancel  context.CancelFunc
	tc      context.Context
	te      *envtest.Environment
	ts      *httptest.Server
)

func init() {
	if err := utils.ConfigureTFReattachProvider(); err != nil {
		panic(err)
	}
}

func TestFunctionality(t *testing.T) {
	RegisterFailHandler(Fail)

	// Ensure verbose log output
	suiteConfig, reporterConfig := GinkgoConfiguration()
	reporterConfig.FullTrace = true
	reporterConfig.Verbose = true

	RunSpecs(t, "Integration Test Suite", suiteConfig, reporterConfig)
}

var _ = BeforeSuite(func() {
	tc = context.Background()
	tc, cancel = context.WithCancel(tc)

	// Start a mock server for the Palette API
	router := mux.NewRouter()
	routes.RegisterProject(router)
	routes.RegisterCloudAccount(router)
	ts = httptest.NewTLSServer(router)

	// Start the test environment
	te = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "..", "package", "crds")},
		ErrorIfCRDPathMissing: true,
	}
	if kcfg := os.Getenv("KUBECONFIG"); kcfg != "" {
		te.UseExistingCluster = utils.Ptr(true)
	}

	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	cfg, err := te.Start()
	Expect(err).ShouldNot(HaveOccurred())

	scheme := scheme.Scheme
	Expect(clusterapis.AddToScheme(scheme)).To(Succeed())

	kclient, err = client.New(cfg, client.Options{Scheme: scheme})
	Expect(err).ShouldNot(HaveOccurred())

	// Generate, save, and configure kubeconfig so in-cluster client lookups succeed
	kubeconfigPath, err := kubeconfigFromRestConfig(cfg)
	Expect(err).ShouldNot(HaveOccurred())

	os.Setenv("KUBECONFIG", kubeconfigPath)
	logf.Log.Info("Kubeconfig", "path", kubeconfigPath)

	initResources()

	go func() {
		defer GinkgoRecover()
		os.Args = initProviderArgs()
		run.Run()
	}()
})

var _ = AfterSuite(func() {
	By("tearing down the mock server")
	ts.Close()

	By("tearing down the test environment")
	cancel()
	err := te.Stop()
	Expect(err).ShouldNot(HaveOccurred(), "failed to tear down the test environment")
})

// initProviderArgs initializes the arguments required to run provider-palette in test mode.
func initProviderArgs() []string {
	v := os.Getenv("TERRAFORM_VERSION")
	if v == "" {
		Fail("TERRAFORM_VERSION environment variable must be set")
	}
	pv := os.Getenv("TERRAFORM_PROVIDER_VERSION")
	if v == "" {
		Fail("TERRAFORM_PROVIDER_VERSION environment variable must be set")
	}
	return []string{
		"run",
		"--debug",
		fmt.Sprintf("--terraform-version=%s", v),
		fmt.Sprintf("--terraform-provider-version=%s", pv),
		"--terraform-provider-source=spectrocloud/spectrocloud",
	}
}

// initResources creates namespace, secret, and provider config resources
// required for provider-palette to function in test mode.
func initResources() {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "crossplane-system",
		},
	}
	if err := kclient.Create(tc, ns); err != nil && !kerrs.IsAlreadyExists(err) {
		Fail(fmt.Sprintf("failed to create namespace: %v", err))
	}

	s := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "palette-creds",
			Namespace: "crossplane-system",
		},
		StringData: map[string]string{
			"credentials": fmt.Sprintf(`{
					"api_key": "dummy",
					"project_name": "Default",
					"host": "%s",
					"ignore_insecure_tls_error": "true",
					"retry_attempts": "1",
					"trace": "true"
				}`,
				ts.Listener.Addr().String(), // hook provider-palette's TF host into the Palette mock server
			),
		},
	}
	if err := kclient.Get(tc, types.NamespacedName{Name: s.Name, Namespace: s.Namespace}, &corev1.Secret{}); err == nil {
		Expect(kclient.Delete(tc, s)).Should(Succeed())
		time.Sleep(interval)
	}
	Expect(kclient.Create(tc, s)).Should(Succeed())

	pc := &clusterv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "provider-palette-config",
			Namespace: "crossplane-system",
		},
		Spec: clusterv1beta1.ProviderConfigSpec{
			Credentials: clusterv1beta1.ProviderCredentials{
				Source: xpv1.CredentialsSource("Secret"),
				CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
					SecretRef: &xpv1.SecretKeySelector{
						SecretReference: xpv1.SecretReference{
							Name:      "palette-creds",
							Namespace: "crossplane-system",
						},
						Key: "credentials",
					},
				},
			},
		},
	}
	if err := kclient.Create(tc, pc); err != nil && !kerrs.IsAlreadyExists(err) {
		Fail(fmt.Sprintf("failed to create ProviderConfig: %v", err))
	}
}
