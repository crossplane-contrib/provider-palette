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
	// Add configurations for resources that should be namespace-scoped
	// These are typically resources that belong to a specific namespace/workspace

	// Example: Application resources that are namespace-scoped
	p.AddResourceConfigurator("spectrocloud_application", func(r *config.Resource) {
		// Ensure this resource is properly scoped to namespaces
		// Add any namespace-specific configurations here
	})

	// Example: Addon deployments that are namespace-scoped
	p.AddResourceConfigurator("spectrocloud_addon_deployment", func(r *config.Resource) {
		// Namespace-specific configurations for addon deployments
	})

	// Add more namespaced resource configurations as needed
}
