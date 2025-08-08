package resources

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	configurePlatformResources(p)
	configureCoreResources(p)
	configureVirtualResources(p)
	configureAccessControlResources(p)
	configureClusterResources(p)
	configureSecurityResources(p)
	configureMonitoringResources(p)
	configureRegistryResources(p)
	configurePrivateCloudGatewayResources(p)
	configureBackupResources(p)
}

// configurePlatformResources configures platform-level resources with customdiff overrides
func configurePlatformResources(p *config.Provider) {
	// Special configuration for platform settings to override customdiff
	p.AddResourceConfigurator("spectrocloud_platform_setting", func(r *config.Resource) {
		// Override customdiff to prevent forced recreation and handle schema validation errors
		if r.TerraformResource != nil && r.TerraformResource.CustomizeDiff != nil {
			r.TerraformResource.CustomizeDiff = nil
		}
	})

	// Special configuration for sso to override customdiff
	p.AddResourceConfigurator("spectrocloud_sso", func(r *config.Resource) {
		// Need fix import(Observe) support for this resource in terraform provider
		// Override customdiff to prevent forced recreation and handle schema validation errors
		if r.TerraformResource != nil && r.TerraformResource.CustomizeDiff != nil {
			r.TerraformResource.CustomizeDiff = nil
		}
	})

	// Special configuration for password policy to override customdiff
	p.AddResourceConfigurator("spectrocloud_password_policy", func(r *config.Resource) {
		// Need fix import(Observe) support for this resource in terraform provider
		// Override customdiff to prevent forced recreation and handle schema validation errors
		if r.TerraformResource != nil && r.TerraformResource.CustomizeDiff != nil {
			r.TerraformResource.CustomizeDiff = nil
		}
	})
}

// configureCoreResources configures core Spectrocloud resources
func configureCoreResources(p *config.Provider) {
	// Spectrocloud core resources
	p.AddResourceConfigurator("spectrocloud_project", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})

	p.AddResourceConfigurator("spectrocloud_workspace", func(r *config.Resource) {
		// Need to add import(Observe) support for this resource in terraform provider
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	// "spectrocloud_macro" - skipped via SkipList (singular only)
	p.AddResourceConfigurator("spectrocloud_macros", func(r *config.Resource) {
		// Need fix import(Observe) support for this resource in terraform provider
	})

	// Application resources
	p.AddResourceConfigurator("spectrocloud_application", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
		r.References["application_profile_uid"] = config.Reference{
			TerraformName: "spectrocloud_application_profile",
		}
		r.References["config.cluster_group_uid"] = config.Reference{
			TerraformName: "spectrocloud_cluster_group",
		}
		r.References["config.cluster_uid"] = config.Reference{
			TerraformName: "spectrocloud_virtual_cluster",
		}
	})

	p.AddResourceConfigurator("spectrocloud_application_profile", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})

	// Cluster profile resources
	p.AddResourceConfigurator("spectrocloud_cluster_profile", func(r *config.Resource) {})

	// Cluster group resources
	p.AddResourceConfigurator("spectrocloud_cluster_group", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
	})

	// Addon resources
	p.AddResourceConfigurator("spectrocloud_addon_deployment", func(r *config.Resource) {
		// Import case is not supported for this resource in terraform provider
		r.UseAsync = true
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
	})

	// Platform resources
	p.AddResourceConfigurator("spectrocloud_resource_limit", func(r *config.Resource) {
		// Need fix import(Observe) support for this resource in terraform provider
	})
}

// configureVirtualResources configures virtual cluster and virtual machine resources
func configureVirtualResources(p *config.Provider) {
	// Virtual cluster resources
	p.AddResourceConfigurator("spectrocloud_virtual_cluster", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
		r.UseAsync = true
		r.References["cluster_group_uid"] = config.Reference{
			TerraformName: "spectrocloud_cluster_group",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
	})

	// Virtual machine resources
	p.AddResourceConfigurator("spectrocloud_virtual_machine", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})

	p.AddResourceConfigurator("spectrocloud_datavolume", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})
}

// configureAccessControlResources configures user, team, and role resources
func configureAccessControlResources(p *config.Provider) {
	p.AddResourceConfigurator("spectrocloud_role", func(r *config.Resource) {
		r.ExternalName = config.TemplatedStringAsIdentifier("", "{{ .external_name }}")
		r.ExternalName.DisableNameInitializer = true
	})

	p.AddResourceConfigurator("spectrocloud_team", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
		r.References["project_role_mapping.id"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["project_role_mapping.roles"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["tenant_role_mapping"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["workspace_role_mapping.id"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["workspace_role_mapping.workspace.roles"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["workspace_role_mapping.workspace.id"] = config.Reference{
			TerraformName: "spectrocloud_workspace",
		}
		r.References["users"] = config.Reference{
			TerraformName: "spectrocloud_user",
		}
	})

	p.AddResourceConfigurator("spectrocloud_user", func(r *config.Resource) {
		r.ExternalName = config.TemplatedStringAsIdentifier("", "{{ .external_name }}")
		r.ExternalName.DisableNameInitializer = true
		r.References["project_role.role_ids"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["project_role.project_id"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["tenant_role"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["workspace_role.project_id"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["workspace_role.workspace.role_ids"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["workspace_role.workspace.id"] = config.Reference{
			TerraformName: "spectrocloud_workspace",
		}
		r.References["resource_role.project_ids"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
		r.References["resource_role.role_ids"] = config.Reference{
			TerraformName: "spectrocloud_role",
		}
		r.References["resource_role.filter_ids"] = config.Reference{
			TerraformName: "spectrocloud_filter",
		}
		r.References["team_ids"] = config.Reference{
			TerraformName: "spectrocloud_team",
		}
	})
}

// configureClusterResources configures cluster resources for all cloud providers
func configureClusterResources(p *config.Provider) {
	// Edge resources
	p.AddResourceConfigurator("spectrocloud_appliance", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})

	p.AddResourceConfigurator("spectrocloud_cluster_edge_native", func(r *config.Resource) {

		r.UseAsync = true
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	// VSphere resources
	p.AddResourceConfigurator("spectrocloud_cloudaccount_vsphere", func(r *config.Resource) {})

	p.AddResourceConfigurator("spectrocloud_cluster_vsphere", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_vsphere",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	p.AddResourceConfigurator("spectrocloud_cluster_edge_vsphere", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	// OpenStack resources
	p.AddResourceConfigurator("spectrocloud_cluster_openstack", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_openstack",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	p.AddResourceConfigurator("spectrocloud_cloudaccount_openstack", func(r *config.Resource) {})

	// MAAS resources
	p.AddResourceConfigurator("spectrocloud_cluster_maas", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_maas",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	p.AddResourceConfigurator("spectrocloud_cloudaccount_maas", func(r *config.Resource) {})

	// Azure resources
	p.AddResourceConfigurator("spectrocloud_cloudaccount_azure", func(r *config.Resource) {})

	p.AddResourceConfigurator("spectrocloud_cluster_azure", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_azure",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	p.AddResourceConfigurator("spectrocloud_cluster_aks", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_azure",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	// GCP resources
	p.AddResourceConfigurator("spectrocloud_cluster_gcp", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_gcp",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	p.AddResourceConfigurator("spectrocloud_cluster_gke", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_gcp",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	p.AddResourceConfigurator("spectrocloud_cloudaccount_gcp", func(r *config.Resource) {
		if s, ok := r.TerraformResource.Schema["gcp_json_credentials"]; ok {
			s.Sensitive = false
			s.Computed = false
			s.Optional = false
		}
	})

	// Custom cloud resources
	p.AddResourceConfigurator("spectrocloud_cluster_custom_cloud", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_custom",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	p.AddResourceConfigurator("spectrocloud_cloudaccount_custom", func(r *config.Resource) {})

	// AWS resources
	p.AddResourceConfigurator("spectrocloud_cloudaccount_aws", func(r *config.Resource) {})

	p.AddResourceConfigurator("spectrocloud_cluster_aws", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_aws",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	p.AddResourceConfigurator("spectrocloud_cluster_eks", func(r *config.Resource) {
		r.UseAsync = true
		r.References["cloud_account_id"] = config.Reference{
			TerraformName: "spectrocloud_cloudaccount_aws",
		}
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})
}

// configureSecurityResources configures security-related resources
func configureSecurityResources(p *config.Provider) {
	p.AddResourceConfigurator("spectrocloud_ssh_key", func(r *config.Resource) {})

	p.AddResourceConfigurator("spectrocloud_registration_token", func(r *config.Resource) {
		r.ExternalName = config.TemplatedStringAsIdentifier("", "{{ .external_name }}")
		r.ExternalName.DisableNameInitializer = true
		r.References["project_uid"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
	})

	p.AddResourceConfigurator("spectrocloud_developer_setting", func(r *config.Resource) {})
}

// configureMonitoringResources configures monitoring and alerting resources
func configureMonitoringResources(p *config.Provider) {
	p.AddResourceConfigurator("spectrocloud_alert", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})

	p.AddResourceConfigurator("spectrocloud_filter", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})
}

// configureRegistryResources configures registry resources
func configureRegistryResources(p *config.Provider) {
	p.AddResourceConfigurator("spectrocloud_registry_helm", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})

	p.AddResourceConfigurator("spectrocloud_registry_oci", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})
}

// configurePrivateCloudGatewayResources configures private cloud gateway resources
func configurePrivateCloudGatewayResources(p *config.Provider) {
	p.AddResourceConfigurator("spectrocloud_privatecloudgateway_dns_map", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
	})

	p.AddResourceConfigurator("spectrocloud_privatecloudgateway_ippool", func(r *config.Resource) {
		// Need add import(Observe)  support for this resource in terraform provider
	})
}

// configureBackupResources configures backup-related resources
func configureBackupResources(p *config.Provider) {
	p.AddResourceConfigurator("spectrocloud_backup_storage_location", func(r *config.Resource) {
		// Need add import(Observe) support for this resource in terraform provider
		r.UseAsync = true
		if gcpBlock, ok := r.TerraformResource.Schema["gcp_storage_config"]; ok {
			if gcpBlockResource, ok := gcpBlock.Elem.(*schema.Resource); ok {
				if s, ok := gcpBlockResource.Schema["gcp_json_credentials"]; ok {
					s.Sensitive = false
					s.Required = true
					s.Optional = false
				}
			}
		}
	})
}
