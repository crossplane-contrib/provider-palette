/*
Copyright 2021 Upbound Inc.
*/

package namespaced

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for namespaced resources.
func Configure(p *config.Provider) {
	// Configure namespaced-specific resources
	configureNamespacedResources(p)
}

// configureNamespacedResources configures resources that should be namespace-scoped
func configureNamespacedResources(p *config.Provider) {
	// Add cluster-specific configurations here if needed
	// Most configurations are handled in common/config.go
}
