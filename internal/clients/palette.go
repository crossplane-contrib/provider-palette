/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/crossplane-contrib/provider-palette/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal palette credentials as JSON"
)

const (
	keyAPIURL      = "host"
	keyProjectName = "project_name"
	keyAPIKey      = "api_key"

	// Palette credentials environment variable names

	envPaletteAPIKey = "PALETTE_API_KEY"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		paletteCreds := map[string]string{}
		if err := json.Unmarshal(data, &paletteCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set provider configuration

		ps.Configuration = map[string]interface{}{
			keyAPIURL:      paletteCreds[keyAPIURL],
			keyProjectName: paletteCreds[keyProjectName],
		}

		if v, ok := paletteCreds[keyProjectName]; ok {
			ps.Configuration[keyProjectName] = v
		}
		if v, ok := paletteCreds[keyAPIKey]; ok {
			ps.Configuration[keyAPIKey] = v
		}

		// set environment variables for sensitive provider configuration
		ps.ClientMetadata = map[string]string{
			envPaletteAPIKey: paletteCreds[keyAPIKey],
		}

		return ps, nil
	}
}
