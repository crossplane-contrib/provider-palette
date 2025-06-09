/*
Copyright 2021 Upbound Inc.
*/

package null

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("null_resource", func(r *ujconfig.Resource) {
		r.Kind = "Resource"
		// And other overrides.
	})

	// List of resources that need external name configuration
	resources := []string{
		"spectrocloud_filter",
		"spectrocloud_macros",
		"spectrocloud_addon_deployment",
		"spectrocloud_alert",
		"spectrocloud_appliance",
		"spectrocloud_cluster_profile_import",
		"spectrocloud_datavolume",
		"spectrocloud_developer_setting",
		"spectrocloud_platform_setting",
		"spectrocloud_password_policy",
		"spectrocloud_privatecloudgateway_dns_map",
		"spectrocloud_resource_limit",
		"spectrocloud_sso",
		"spectrocloud_user",
	}

	// Configure each resource
	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *ujconfig.Resource) {
			// Prevent Crossplane from injecting external-name into Terraform JSON
			r.ExternalName = ujconfig.IdentifierFromProvider
		})
	}
}
