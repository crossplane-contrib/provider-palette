/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-palette/config/null"
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
		config.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		config.WithSkipList([]string{
			"import",
		}),
	)

	for _, configure := range []func(provider *config.Provider){
		// add custom config functions
		null.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
