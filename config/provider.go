/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-palette/config/resources"
)

const (
	resourcePrefix = "palette"
	modulePath     = "github.com/crossplane-contrib/provider-palette"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *config.Provider {
	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		config.WithShortName("palette"),
		config.WithRootGroup("palette.crossplane.io"),
		config.WithFeaturesPackage("internal/features"),
		config.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		config.WithSkipList([]string{
			"spectrocloud_cluster_profile_import", // Specific resource to skip
			"^spectrocloud_macro$",                // Only skip exact singular match, keep spectrocloud_macros
		}),
	)

	for _, configure := range []func(provider *config.Provider){
		// add custom config functions
		resources.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
