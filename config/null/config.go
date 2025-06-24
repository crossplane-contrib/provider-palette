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
	// Special configuration for platform settings to override customdiff
	p.AddResourceConfigurator("spectrocloud_platform_setting", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		// Override customdiff to prevent forced recreation and handle schema validation errors
		if r.TerraformResource != nil && r.TerraformResource.CustomizeDiff != nil {
			r.TerraformResource.CustomizeDiff = nil
		}
		r.UseAsync = true
	})

	// Special configuration for platform settings to override customdiff
	p.AddResourceConfigurator("spectrocloud_sso", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		// Override customdiff to prevent forced recreation and handle schema validation errors
		if r.TerraformResource != nil && r.TerraformResource.CustomizeDiff != nil {
			r.TerraformResource.CustomizeDiff = nil
		}
		r.UseAsync = true
	})
	// Special configuration for platform settings to override customdiff
	p.AddResourceConfigurator("spectrocloud_password_policy", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		// Override customdiff to prevent forced recreation and handle schema validation errors
		if r.TerraformResource != nil && r.TerraformResource.CustomizeDiff != nil {
			r.TerraformResource.CustomizeDiff = nil
		}
		r.UseAsync = true
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
		"spectrocloud_privatecloudgateway_dns_map",
		"spectrocloud_resource_limit",
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
