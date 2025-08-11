/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-palette/config/cluster"
	"github.com/crossplane-contrib/provider-palette/config/common"
	"github.com/crossplane-contrib/provider-palette/config/namespaced"
)

const (
	resourcePrefix = "palette"
	modulePath     = "github.com/crossplane-contrib/provider-palette"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

var commonSkipList = []string{
	"spectrocloud_cluster_profile_import", // Specific resource to skip
	"^spectrocloud_macro$",                // Only skip exact singular match, keep spectrocloud_macros
}

// GetProvider returns provider configuration for cluster-scoped resources
func GetProvider() *config.Provider {
	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		config.WithShortName("palette"),
		config.WithRootGroup("palette.crossplane.io"),
		config.WithFeaturesPackage("internal/features"),

		config.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		config.WithSkipList(commonSkipList),
	)

	for _, configure := range []func(provider *config.Provider){
		// add custom config functions
		cluster.Configure,
		common.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns provider configuration for namespaced resources
func GetProviderNamespaced() *config.Provider {
	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		config.WithShortName("palette"),
		config.WithRootGroup("palette.m.crossplane.io"), // Note: .m. for managed/namespaced
		config.WithFeaturesPackage("internal/features"),

		config.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		config.WithSkipList(commonSkipList),
	)

	for _, configure := range []func(provider *config.Provider){
		// add custom config functions
		namespaced.Configure,
		common.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
