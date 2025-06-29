package resources

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// Special configuration for platform settings to override customdiff
	p.AddResourceConfigurator("spectrocloud_platform_setting", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "PlatformSetting"
		// Override customdiff to prevent forced recreation and handle schema validation errors
		if r.TerraformResource != nil && r.TerraformResource.CustomizeDiff != nil {
			r.TerraformResource.CustomizeDiff = nil
		}
	})

	// Special configuration for sso to override customdiff
	p.AddResourceConfigurator("spectrocloud_sso", func(r *config.Resource) {
		r.ShortGroup = "security"
		r.Kind = "SSO"
		// Override customdiff to prevent forced recreation and handle schema validation errors
		if r.TerraformResource != nil && r.TerraformResource.CustomizeDiff != nil {
			r.TerraformResource.CustomizeDiff = nil
		}
	})

	// Special configuration for password policy to override customdiff
	p.AddResourceConfigurator("spectrocloud_password_policy", func(r *config.Resource) {
		r.ShortGroup = "security"
		r.Kind = "PasswordPolicy"
		// Override customdiff to prevent forced recreation and handle schema validation errors
		if r.TerraformResource != nil && r.TerraformResource.CustomizeDiff != nil {
			r.TerraformResource.CustomizeDiff = nil
		}
	})

	// Spectrocloud core resources
	p.AddResourceConfigurator("spectrocloud_project", func(r *config.Resource) {
		r.ShortGroup = "spectrocloud"
		r.Kind = "Project"
	})

	p.AddResourceConfigurator("spectrocloud_workspace", func(r *config.Resource) {
		r.ShortGroup = "spectrocloud"
		r.Kind = "Workspace"
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	p.AddResourceConfigurator("spectrocloud_macro", func(r *config.Resource) {
		r.ShortGroup = "spectrocloud"
		r.Kind = "Macro"
	})

	p.AddResourceConfigurator("spectrocloud_macros", func(r *config.Resource) {
		r.ShortGroup = "spectrocloud"
		r.Kind = "Macros"
	})

	// Virtual cluster resources
	p.AddResourceConfigurator("spectrocloud_virtual_cluster", func(r *config.Resource) {
		r.ShortGroup = "virtualcluster"
		r.Kind = "VirtualCluster"
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

	// Access control resources
	p.AddResourceConfigurator("spectrocloud_role", func(r *config.Resource) {
		r.ShortGroup = "accesscontrol"
		r.Kind = "Role"
	})

	p.AddResourceConfigurator("spectrocloud_team", func(r *config.Resource) {
		r.ShortGroup = "accesscontrol"
		r.Kind = "Team"
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
		r.ShortGroup = "accesscontrol"
		r.Kind = "User"
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

	// Cluster group resources
	p.AddResourceConfigurator("spectrocloud_cluster_group", func(r *config.Resource) {
		r.ShortGroup = "clustergroup"
		r.Kind = "ClusterGroup"
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
	})

	// Application resources
	p.AddResourceConfigurator("spectrocloud_application", func(r *config.Resource) {
		r.ShortGroup = "application"
		r.Kind = "Application"
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
		r.ShortGroup = "application"
		r.Kind = "ApplicationProfile"
	})

	// Cluster profile resources
	p.AddResourceConfigurator("spectrocloud_cluster_profile", func(r *config.Resource) {
		r.ShortGroup = "clusterprofile"
		r.Kind = "ClusterProfile"
	})

	p.AddResourceConfigurator("spectrocloud_cluster_profile_import", func(r *config.Resource) {
		r.ShortGroup = "clusterprofile"
		r.Kind = "ClusterProfileImport"
	})

	// Addon resources
	p.AddResourceConfigurator("spectrocloud_addon_deployment", func(r *config.Resource) {
		r.ShortGroup = "spectrocloud"
		r.Kind = "AddonDeployment"
		r.UseAsync = true
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
	})

	// Virtual machine resources
	p.AddResourceConfigurator("spectrocloud_virtual_machine", func(r *config.Resource) {
		r.ShortGroup = "virtualmachine"
		r.Kind = "VirtualMachine"
	})

	p.AddResourceConfigurator("spectrocloud_datavolume", func(r *config.Resource) {
		r.ShortGroup = "virtualmachine"
		r.Kind = "DataVolume"
	})

	// Edge resources
	p.AddResourceConfigurator("spectrocloud_appliance", func(r *config.Resource) {
		r.ShortGroup = "edge"
		r.Kind = "Appliance"
	})

	p.AddResourceConfigurator("spectrocloud_cluster_edge_native", func(r *config.Resource) {
		r.ShortGroup = "edge"
		r.Kind = "EdgeCluster"
		r.UseAsync = true
		r.References["cluster_profile.id"] = config.Reference{
			TerraformName: "spectrocloud_cluster_profile",
		}
		r.References["backup_policy.backup_location_id"] = config.Reference{
			TerraformName: "spectrocloud_backup_storage_location",
		}
	})

	// VSphere resources
	p.AddResourceConfigurator("spectrocloud_cloudaccount_vsphere", func(r *config.Resource) {
		r.ShortGroup = "vsphere"
		r.Kind = "VSphereCloudAccount"
	})

	p.AddResourceConfigurator("spectrocloud_cluster_vsphere", func(r *config.Resource) {
		r.ShortGroup = "vsphere"
		r.Kind = "VSphereCluster"
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
		r.ShortGroup = "vsphere"
		r.Kind = "EdgeVSphereCluster"
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
		r.ShortGroup = "openstack"
		r.Kind = "OpenstackCluster"
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

	p.AddResourceConfigurator("spectrocloud_cloudaccount_openstack", func(r *config.Resource) {
		r.ShortGroup = "openstack"
		r.Kind = "OpenstackCloudAccount"
	})

	// MAAS resources
	p.AddResourceConfigurator("spectrocloud_cluster_maas", func(r *config.Resource) {
		r.ShortGroup = "maas"
		r.Kind = "MAASCluster"
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

	p.AddResourceConfigurator("spectrocloud_cloudaccount_maas", func(r *config.Resource) {
		r.ShortGroup = "maas"
		r.Kind = "MAASCloudAccount"
	})

	// Azure resources
	p.AddResourceConfigurator("spectrocloud_cloudaccount_azure", func(r *config.Resource) {
		r.ShortGroup = "azure"
		r.Kind = "AzureCloudAccount"
	})

	p.AddResourceConfigurator("spectrocloud_cluster_azure", func(r *config.Resource) {
		r.ShortGroup = "azure"
		r.Kind = "AzureCluster"
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
		r.ShortGroup = "azure"
		r.Kind = "AKSCluster"
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
		r.ShortGroup = "gcp"
		r.Kind = "GCPCluster"
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
		r.ShortGroup = "gcp"
		r.Kind = "GKECluster"
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
		r.ShortGroup = "gcp"
		r.Kind = "GCPCloudAccount"
	})

	// Custom cloud resources
	p.AddResourceConfigurator("spectrocloud_cluster_custom_cloud", func(r *config.Resource) {
		r.ShortGroup = "customcloud"
		r.Kind = "CustomCluster"
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

	p.AddResourceConfigurator("spectrocloud_cloudaccount_custom", func(r *config.Resource) {
		r.ShortGroup = "customcloud"
		r.Kind = "CustomCloudAccount"
	})

	// AWS resources
	p.AddResourceConfigurator("spectrocloud_cloudaccount_aws", func(r *config.Resource) {
		r.ShortGroup = "aws"
		r.Kind = "AWSCloudAccount"
	})

	p.AddResourceConfigurator("spectrocloud_cluster_aws", func(r *config.Resource) {
		r.ShortGroup = "aws"
		r.Kind = "AWSCluster"
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
		r.ShortGroup = "aws"
		r.Kind = "EKSCluster"
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

	// Security resources
	p.AddResourceConfigurator("spectrocloud_ssh_key", func(r *config.Resource) {
		r.ShortGroup = "security"
		r.Kind = "SSHKey"
	})

	p.AddResourceConfigurator("spectrocloud_registration_token", func(r *config.Resource) {
		r.ShortGroup = "security"
		r.Kind = "RegistrationToken"
		r.References["project_uid"] = config.Reference{
			TerraformName: "spectrocloud_project",
		}
	})

	p.AddResourceConfigurator("spectrocloud_developer_setting", func(r *config.Resource) {
		r.ShortGroup = "security"
		r.Kind = "DeveloperSetting"
	})

	// Monitoring resources
	p.AddResourceConfigurator("spectrocloud_alert", func(r *config.Resource) {
		r.ShortGroup = "monitoring"
		r.Kind = "Alert"
	})

	p.AddResourceConfigurator("spectrocloud_filter", func(r *config.Resource) {
		r.ShortGroup = "monitoring"
		r.Kind = "Filter"
	})

	// Platform resources
	p.AddResourceConfigurator("spectrocloud_resource_limit", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "ResourceLimit"
	})

	// Registry resources
	p.AddResourceConfigurator("spectrocloud_registry_helm", func(r *config.Resource) {
		r.ShortGroup = "registry"
		r.Kind = "HelmRegistry"
	})

	p.AddResourceConfigurator("spectrocloud_registry_oci", func(r *config.Resource) {
		r.ShortGroup = "registry"
		r.Kind = "OCIRegistry"
	})

	// Private cloud gateway resources
	p.AddResourceConfigurator("spectrocloud_privatecloudgateway_dns_map", func(r *config.Resource) {
		r.ShortGroup = "privatecloudgateway"
		r.Kind = "DNSMap"
	})

	p.AddResourceConfigurator("spectrocloud_privatecloudgateway_ippool", func(r *config.Resource) {
		r.ShortGroup = "privatecloudgateway"
		r.Kind = "IPPool"
	})

	// Backup resources
	p.AddResourceConfigurator("spectrocloud_backup_storage_location", func(r *config.Resource) {
		r.ShortGroup = "backup"
		r.Kind = "BackupStorageLocation"
		r.UseAsync = true
	})
}
