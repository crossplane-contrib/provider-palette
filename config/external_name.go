/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// FormattedIdentifierFromProvider returns external name configuration for resources that use
// composite identifiers. Similar to AWS provider pattern - handles bidirectional conversion between
// user-friendly name and composite identifier (name:context) for Spectrocloud resources.
func FormattedIdentifierFromProvider() config.ExternalName {
	e := config.IdentifierFromProvider // Use IdentifierFromProvider to keep name field in spec

	// CREATE/IMPORT: use appropriate identifier + context → composite ID for Terraform
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
		var identifier string

		// During CREATE: externalName is empty, use name from parameters
		// During IMPORT: externalName is set from annotation, use it
		if externalName != "" {
			identifier = externalName
		} else {
			// Fallback to name parameter for CREATE operations
			name, ok := parameters["name"].(string)
			if !ok || name == "" {
				return "", errors.New("name parameter is required when external name is not provided")
			}
			identifier = name
		}

		return identifier, nil
	}

	// IMPORT: terraform state ID → external name
	// The terraform provider stores composite ID in format "id:context"
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id attribute missing from terraform state")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("id must be a string")
		}

		// Split composite ID and return only the ID part
		// Example: "687f86db08e0e7ea51677fa3:project" → "687f86db08e0e7ea51677fa3"
		parts := strings.Split(idStr, ":")
		if len(parts) >= 2 {
			return parts[0], nil // Return the ID part
		}
		// Fallback for simple ID format (if any)
		return idStr, nil
	}

	return e
}

func FormattedIdentifierPCG() config.ExternalName {
	e := config.IdentifierFromProvider // Use IdentifierFromProvider to keep name field in spec

	// CREATE/IMPORT: use appropriate identifier + context → composite ID for Terraform
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
		var identifier string

		// During CREATE: externalName is empty, use name from parameters
		// During IMPORT: externalName is set from annotation, use it
		if externalName != "" {
			identifier = externalName
		} else {
			// Fallback to name parameter for CREATE operations
			name, ok := parameters["name"].(string)
			if !ok || name == "" {
				return "", errors.New("name parameter is required when external name is not provided")
			}
			identifier = name
		}

		return identifier, nil
	}

	// IMPORT: terraform state ID → external name
	// The terraform provider stores composite ID in format "id:context"
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id attribute missing from terraform state")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("id must be a string")
		}

		// Split composite ID and return only the ID part
		// Example: "687f86db08e0e7ea51677fa3:d34236db08e0e7ea5167233sw" → "d34236db08e0e7ea5167233sw"
		parts := strings.Split(idStr, ":")
		if len(parts) >= 2 {
			return parts[1], nil // Return the ID part
		}
		// Fallback for simple ID format (if any)
		return idStr, nil
	}

	return e
}

// platformSetting returns external name configuration for spectrocloud_platform_setting
// Handles the special "platformsetting-" prefix pattern: "platformsetting-{external_name}:{context}"
func platformSetting() config.ExternalName {
	e := config.IdentifierFromProvider // Use IdentifierFromProvider to keep name field in spec

	// CREATE/IMPORT: use appropriate identifier + context → composite ID with prefix for Terraform
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
		var identifier string

		// During CREATE: externalName is empty, use name from parameters
		// During IMPORT: externalName is set from annotation, use it
		if externalName != "" {
			identifier = externalName
		} else {
			// Fallback to name parameter for CREATE operations
			name, ok := parameters["name"].(string)
			if !ok || name == "" {
				return "", errors.New("name parameter is required when external name is not provided")
			}
			identifier = name
		}

		// Add the "platformsetting-" prefix
		return fmt.Sprintf("platformsetting-%s", identifier), nil
	}

	// IMPORT: terraform state ID → external name
	// The terraform provider stores composite ID with prefix: "platformsetting-name:context"
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id attribute missing from terraform state")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("id must be a string")
		}

		// Split composite ID and extract name part from prefix
		// Example: "platformsetting-myconfig:tenant" → "myconfig"
		parts := strings.Split(idStr, ":")
		if len(parts) >= 2 {
			namePart := parts[0]
			// Remove the "platformsetting-" prefix if present
			if strings.HasPrefix(namePart, "platformsetting-") {
				return strings.TrimPrefix(namePart, "platformsetting-"), nil
			}
			return namePart, nil
		}
		// Fallback for simple ID format (if any)
		return idStr, nil
	}

	return e
}

// cloudAccountCustom returns external name configuration for spectrocloud_cloudaccount_custom
// Handles the special 3-part pattern: "id:context:type" (e.g., "account123:project:nutanix")
func cloudAccountCustom() config.ExternalName {
	e := config.IdentifierFromProvider // Use IdentifierFromProvider to keep name field in spec

	// CREATE/IMPORT: use appropriate identifier + context + cloud → 3-part composite ID for Terraform
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
		var identifier string

		// During CREATE: externalName is empty, use name from parameters
		// During IMPORT: externalName is set from annotation, use it
		if externalName != "" {
			identifier = externalName
		} else {
			// Fallback to name parameter for CREATE operations
			name, ok := parameters["name"].(string)
			if !ok || name == "" {
				return "", errors.New("name parameter is required when external name is not provided")
			}
			identifier = name
		}

		return identifier, nil
	}

	// IMPORT: terraform state ID → external name
	// The terraform provider returns 3-part ID format: "id:context:type"
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id attribute missing from terraform state")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("id must be a string")
		}

		// Parse 3-part format: "id:context:type" and extract just the ID part
		parts := strings.Split(idStr, ":")
		if len(parts) == 3 {
			// Return just the ID part (first part before first ":")
			return parts[0], nil
		}

		// Fallback: if not 3-part format, return the whole ID
		return idStr, nil
	}

	return e
}

// clusterCustomCloud returns external name configuration for spectrocloud_cluster_custom_cloud
// Handles the special 3-part pattern: "id:context:type" (e.g., "cluster456:project:nutanix")
func clusterCustomCloud() config.ExternalName {
	e := config.IdentifierFromProvider // Use IdentifierFromProvider to keep name field in spec

	// CREATE/IMPORT: use appropriate identifier + context + cloud → 3-part composite ID for Terraform
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
		var identifier string

		// During CREATE: externalName is empty, use name from parameters
		// During IMPORT: externalName is set from annotation, use it
		if externalName != "" {
			identifier = externalName
		} else {
			// Fallback to name parameter for CREATE operations
			name, ok := parameters["name"].(string)
			if !ok || name == "" {
				return "", errors.New("name parameter is required when external name is not provided")
			}
			identifier = name
		}

		return identifier, nil
	}

	// IMPORT: terraform state ID → external name
	// The terraform provider returns 3-part ID format: "id:context:type"
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id attribute missing from terraform state")
		}

		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("id must be a string")
		}

		// Parse 3-part format: "id:context:type" and extract just the ID part
		parts := strings.Split(idStr, ":")
		if len(parts) == 3 {
			// Return just the ID part (first part before first ":")
			return parts[0], nil
		}

		// Fallback: if not 3-part format, return the whole ID
		return idStr, nil
	}

	return e
}

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Spectro Cloud resources - most use provider-generated IDs
	"spectrocloud_addon_deployment":              config.IdentifierFromProvider, // Observe is not support operational resource
	"spectrocloud_alert":                         config.IdentifierFromProvider,
	"spectrocloud_appliance":                     config.IdentifierFromProvider,
	"spectrocloud_application":                   config.IdentifierFromProvider,
	"spectrocloud_application_profile":           config.IdentifierFromProvider,
	"spectrocloud_backup_storage_location":       FormattedIdentifierFromProvider(),
	"spectrocloud_cloudaccount_apachecloudstack": FormattedIdentifierFromProvider(),
	"spectrocloud_cloudaccount_aws":              FormattedIdentifierFromProvider(),
	"spectrocloud_cloudaccount_azure":            FormattedIdentifierFromProvider(),
	"spectrocloud_cloudaccount_custom":           cloudAccountCustom(),
	"spectrocloud_cloudaccount_gcp":              FormattedIdentifierFromProvider(),
	"spectrocloud_cloudaccount_maas":             FormattedIdentifierFromProvider(),
	"spectrocloud_cloudaccount_openstack":        FormattedIdentifierFromProvider(),
	"spectrocloud_cloudaccount_vsphere":          FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_aks":                   FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_apachecloudstack":      FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_aws":                   FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_azure":                 FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_custom_cloud":          clusterCustomCloud(),
	"spectrocloud_cluster_edge_native":           FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_edge_vsphere":          FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_eks":                   FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_gcp":                   FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_gke":                   FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_group":                 FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_maas":                  FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_openstack":             FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_profile":               FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_config_template":       FormattedIdentifierFromProvider(),
	"spectrocloud_cluster_config_policy":         FormattedIdentifierFromProvider(),
	// "spectrocloud_cluster_profile_import" - skipped via SkipList
	"spectrocloud_cluster_vsphere":             FormattedIdentifierFromProvider(),
	"spectrocloud_datavolume":                  config.IdentifierFromProvider, // Observe is not support
	"spectrocloud_developer_setting":           config.IdentifierFromProvider, // Required fix in terrafonfm provider currently observe only is not supported
	"spectrocloud_filter":                      config.IdentifierFromProvider,
	"spectrocloud_macros":                      config.IdentifierFromProvider,
	"spectrocloud_password_policy":             config.IdentifierFromProvider,
	"spectrocloud_platform_setting":            platformSetting(),
	"spectrocloud_privatecloudgateway_dns_map": FormattedIdentifierPCG(),
	"spectrocloud_privatecloudgateway_ippool":  FormattedIdentifierPCG(),
	"spectrocloud_project":                     config.IdentifierFromProvider,
	"spectrocloud_registration_token":          config.IdentifierFromProvider,
	"spectrocloud_registry_helm":               config.IdentifierFromProvider,
	"spectrocloud_registry_oci":                config.IdentifierFromProvider,
	"spectrocloud_resource_limit":              config.IdentifierFromProvider, // Required fix in terrafonfm provider currently observe only is not supported
	"spectrocloud_role":                        config.IdentifierFromProvider,
	"spectrocloud_ssh_key":                     FormattedIdentifierFromProvider(),
	"spectrocloud_sso":                         config.IdentifierFromProvider, // Required fix in terrafonfm provider currently observe only is not supported
	"spectrocloud_team":                        config.IdentifierFromProvider,
	"spectrocloud_user":                        config.IdentifierFromProvider,
	"spectrocloud_virtual_cluster":             config.IdentifierFromProvider,
	"spectrocloud_virtual_machine":             config.IdentifierFromProvider, // Not added support in provider
	"spectrocloud_workspace":                   config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		} else {
			// Default to NameAsExternalName, unless overridden via ExternalNameConfigs.
			r.ExternalName = config.NameAsIdentifier
			r.ExternalName.GetExternalNameFn = NameAsExternalName
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

// NameAsExternalName returns tfstate["name"] as the resource's external name.
func NameAsExternalName(tfstate map[string]any) (string, error) {
	if id, ok := tfstate["name"].(string); ok && id != "" {
		return id, nil
	}
	return "", errors.New("cannot find name in tfstate")
}
