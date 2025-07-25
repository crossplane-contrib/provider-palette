/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BackupPolicyInitParameters struct {

	// (String) The ID of the backup location to use for the backup.
	// The ID of the backup location to use for the backup.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-palette/apis/backup/v1alpha1.StorageLocation
	BackupLocationID *string `json:"backupLocationId,omitempty" tf:"backup_location_id,omitempty"`

	// Reference to a StorageLocation in backup to populate backupLocationId.
	// +kubebuilder:validation:Optional
	BackupLocationIDRef *v1.Reference `json:"backupLocationIdRef,omitempty" tf:"-"`

	// Selector for a StorageLocation in backup to populate backupLocationId.
	// +kubebuilder:validation:Optional
	BackupLocationIDSelector *v1.Selector `json:"backupLocationIdSelector,omitempty" tf:"-"`

	// (Set of String) The list of cluster UIDs to include in the backup. If include_all_clusters is set to true, then all clusters will be included.
	// The list of cluster UIDs to include in the backup. If `include_all_clusters` is set to `true`, then all clusters will be included.
	// +listType=set
	ClusterUids []*string `json:"clusterUids,omitempty" tf:"cluster_uids,omitempty"`

	// (Number) The number of hours after which the backup will be deleted. For example, if the expiry is set to 24, the backup will be deleted after 24 hours.
	// The number of hours after which the backup will be deleted. For example, if the expiry is set to 24, the backup will be deleted after 24 hours.
	ExpiryInHour *float64 `json:"expiryInHour,omitempty" tf:"expiry_in_hour,omitempty"`

	// (Boolean) Whether to include all clusters in the backup. If set to false, only the clusters specified in cluster_uids will be included.
	// Whether to include all clusters in the backup. If set to false, only the clusters specified in `cluster_uids` will be included.
	IncludeAllClusters *bool `json:"includeAllClusters,omitempty" tf:"include_all_clusters,omitempty"`

	// (Boolean) Indicates whether to include cluster resources in the backup. If set to false, only the cluster configuration and disks will be backed up. (Note: Starting with Palette version 4.6, the include_cluster_resources attribute will be deprecated, and a new attribute, include_cluster_resources_mode, will be introduced.)
	// Indicates whether to include cluster resources in the backup. If set to false, only the cluster configuration and disks will be backed up. (Note: Starting with Palette version 4.6, the include_cluster_resources attribute will be deprecated, and a new attribute, include_cluster_resources_mode, will be introduced.)
	IncludeClusterResources *bool `json:"includeClusterResources,omitempty" tf:"include_cluster_resources,omitempty"`

	// (String) Specifies whether to include the cluster resources in the backup. Supported values are always, never, and auto.
	// Specifies whether to include the cluster resources in the backup. Supported values are `always`, `never`, and `auto`.
	IncludeClusterResourcesMode *string `json:"includeClusterResourcesMode,omitempty" tf:"include_cluster_resources_mode,omitempty"`

	// (Boolean) Whether to include the disks in the backup. If set to false, only the cluster configuration will be backed up.
	// Whether to include the disks in the backup. If set to false, only the cluster configuration will be backed up.
	IncludeDisks *bool `json:"includeDisks,omitempty" tf:"include_disks,omitempty"`

	// (Block List) The namespaces for the cluster. (see below for nested schema)
	// The list of Kubernetes namespaces to include in the backup. If not specified, all namespaces will be included.
	// +listType=set
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// -.
	// Prefix for the backup name. The backup name will be of the format <prefix>-<cluster-name>-<timestamp>.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// (String) The schedule for the backup. The schedule is specified in cron format. For example, to run the backup every day at 1:00 AM, the schedule should be set to 0 1 * * *.
	// The schedule for the backup. The schedule is specified in cron format. For example, to run the backup every day at 1:00 AM, the schedule should be set to `0 1 * * *`.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type BackupPolicyObservation struct {

	// (String) The ID of the backup location to use for the backup.
	// The ID of the backup location to use for the backup.
	BackupLocationID *string `json:"backupLocationId,omitempty" tf:"backup_location_id,omitempty"`

	// (Set of String) The list of cluster UIDs to include in the backup. If include_all_clusters is set to true, then all clusters will be included.
	// The list of cluster UIDs to include in the backup. If `include_all_clusters` is set to `true`, then all clusters will be included.
	// +listType=set
	ClusterUids []*string `json:"clusterUids,omitempty" tf:"cluster_uids,omitempty"`

	// (Number) The number of hours after which the backup will be deleted. For example, if the expiry is set to 24, the backup will be deleted after 24 hours.
	// The number of hours after which the backup will be deleted. For example, if the expiry is set to 24, the backup will be deleted after 24 hours.
	ExpiryInHour *float64 `json:"expiryInHour,omitempty" tf:"expiry_in_hour,omitempty"`

	// (Boolean) Whether to include all clusters in the backup. If set to false, only the clusters specified in cluster_uids will be included.
	// Whether to include all clusters in the backup. If set to false, only the clusters specified in `cluster_uids` will be included.
	IncludeAllClusters *bool `json:"includeAllClusters,omitempty" tf:"include_all_clusters,omitempty"`

	// (Boolean) Indicates whether to include cluster resources in the backup. If set to false, only the cluster configuration and disks will be backed up. (Note: Starting with Palette version 4.6, the include_cluster_resources attribute will be deprecated, and a new attribute, include_cluster_resources_mode, will be introduced.)
	// Indicates whether to include cluster resources in the backup. If set to false, only the cluster configuration and disks will be backed up. (Note: Starting with Palette version 4.6, the include_cluster_resources attribute will be deprecated, and a new attribute, include_cluster_resources_mode, will be introduced.)
	IncludeClusterResources *bool `json:"includeClusterResources,omitempty" tf:"include_cluster_resources,omitempty"`

	// (String) Specifies whether to include the cluster resources in the backup. Supported values are always, never, and auto.
	// Specifies whether to include the cluster resources in the backup. Supported values are `always`, `never`, and `auto`.
	IncludeClusterResourcesMode *string `json:"includeClusterResourcesMode,omitempty" tf:"include_cluster_resources_mode,omitempty"`

	// (Boolean) Whether to include the disks in the backup. If set to false, only the cluster configuration will be backed up.
	// Whether to include the disks in the backup. If set to false, only the cluster configuration will be backed up.
	IncludeDisks *bool `json:"includeDisks,omitempty" tf:"include_disks,omitempty"`

	// (Block List) The namespaces for the cluster. (see below for nested schema)
	// The list of Kubernetes namespaces to include in the backup. If not specified, all namespaces will be included.
	// +listType=set
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// -.
	// Prefix for the backup name. The backup name will be of the format <prefix>-<cluster-name>-<timestamp>.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// (String) The schedule for the backup. The schedule is specified in cron format. For example, to run the backup every day at 1:00 AM, the schedule should be set to 0 1 * * *.
	// The schedule for the backup. The schedule is specified in cron format. For example, to run the backup every day at 1:00 AM, the schedule should be set to `0 1 * * *`.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type BackupPolicyParameters struct {

	// (String) The ID of the backup location to use for the backup.
	// The ID of the backup location to use for the backup.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-palette/apis/backup/v1alpha1.StorageLocation
	// +kubebuilder:validation:Optional
	BackupLocationID *string `json:"backupLocationId,omitempty" tf:"backup_location_id,omitempty"`

	// Reference to a StorageLocation in backup to populate backupLocationId.
	// +kubebuilder:validation:Optional
	BackupLocationIDRef *v1.Reference `json:"backupLocationIdRef,omitempty" tf:"-"`

	// Selector for a StorageLocation in backup to populate backupLocationId.
	// +kubebuilder:validation:Optional
	BackupLocationIDSelector *v1.Selector `json:"backupLocationIdSelector,omitempty" tf:"-"`

	// (Set of String) The list of cluster UIDs to include in the backup. If include_all_clusters is set to true, then all clusters will be included.
	// The list of cluster UIDs to include in the backup. If `include_all_clusters` is set to `true`, then all clusters will be included.
	// +kubebuilder:validation:Optional
	// +listType=set
	ClusterUids []*string `json:"clusterUids,omitempty" tf:"cluster_uids,omitempty"`

	// (Number) The number of hours after which the backup will be deleted. For example, if the expiry is set to 24, the backup will be deleted after 24 hours.
	// The number of hours after which the backup will be deleted. For example, if the expiry is set to 24, the backup will be deleted after 24 hours.
	// +kubebuilder:validation:Optional
	ExpiryInHour *float64 `json:"expiryInHour" tf:"expiry_in_hour,omitempty"`

	// (Boolean) Whether to include all clusters in the backup. If set to false, only the clusters specified in cluster_uids will be included.
	// Whether to include all clusters in the backup. If set to false, only the clusters specified in `cluster_uids` will be included.
	// +kubebuilder:validation:Optional
	IncludeAllClusters *bool `json:"includeAllClusters,omitempty" tf:"include_all_clusters,omitempty"`

	// (Boolean) Indicates whether to include cluster resources in the backup. If set to false, only the cluster configuration and disks will be backed up. (Note: Starting with Palette version 4.6, the include_cluster_resources attribute will be deprecated, and a new attribute, include_cluster_resources_mode, will be introduced.)
	// Indicates whether to include cluster resources in the backup. If set to false, only the cluster configuration and disks will be backed up. (Note: Starting with Palette version 4.6, the include_cluster_resources attribute will be deprecated, and a new attribute, include_cluster_resources_mode, will be introduced.)
	// +kubebuilder:validation:Optional
	IncludeClusterResources *bool `json:"includeClusterResources,omitempty" tf:"include_cluster_resources,omitempty"`

	// (String) Specifies whether to include the cluster resources in the backup. Supported values are always, never, and auto.
	// Specifies whether to include the cluster resources in the backup. Supported values are `always`, `never`, and `auto`.
	// +kubebuilder:validation:Optional
	IncludeClusterResourcesMode *string `json:"includeClusterResourcesMode,omitempty" tf:"include_cluster_resources_mode,omitempty"`

	// (Boolean) Whether to include the disks in the backup. If set to false, only the cluster configuration will be backed up.
	// Whether to include the disks in the backup. If set to false, only the cluster configuration will be backed up.
	// +kubebuilder:validation:Optional
	IncludeDisks *bool `json:"includeDisks,omitempty" tf:"include_disks,omitempty"`

	// (Block List) The namespaces for the cluster. (see below for nested schema)
	// The list of Kubernetes namespaces to include in the backup. If not specified, all namespaces will be included.
	// +kubebuilder:validation:Optional
	// +listType=set
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// -.
	// Prefix for the backup name. The backup name will be of the format <prefix>-<cluster-name>-<timestamp>.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// (String) The schedule for the backup. The schedule is specified in cron format. For example, to run the backup every day at 1:00 AM, the schedule should be set to 0 1 * * *.
	// The schedule for the backup. The schedule is specified in cron format. For example, to run the backup every day at 1:00 AM, the schedule should be set to `0 1 * * *`.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type ClusterRbacBindingInitParameters struct {

	// (String) The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (Map of String) The role of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// The role of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// +mapType=granular
	Role map[string]*string `json:"role,omitempty" tf:"role,omitempty"`

	// (Block List) (see below for nested schema)
	Subjects []SubjectsInitParameters `json:"subjects,omitempty" tf:"subjects,omitempty"`

	// (String) The type of the RBAC binding. Can be one of the following values: RoleBinding, or ClusterRoleBinding.
	// The type of the RBAC binding. Can be one of the following values: `RoleBinding`, or `ClusterRoleBinding`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClusterRbacBindingObservation struct {

	// (String) The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (Map of String) The role of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// The role of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// +mapType=granular
	Role map[string]*string `json:"role,omitempty" tf:"role,omitempty"`

	// (Block List) (see below for nested schema)
	Subjects []SubjectsObservation `json:"subjects,omitempty" tf:"subjects,omitempty"`

	// (String) The type of the RBAC binding. Can be one of the following values: RoleBinding, or ClusterRoleBinding.
	// The type of the RBAC binding. Can be one of the following values: `RoleBinding`, or `ClusterRoleBinding`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClusterRbacBindingParameters struct {

	// (String) The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (Map of String) The role of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// The role of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Role map[string]*string `json:"role,omitempty" tf:"role,omitempty"`

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Subjects []SubjectsParameters `json:"subjects,omitempty" tf:"subjects,omitempty"`

	// (String) The type of the RBAC binding. Can be one of the following values: RoleBinding, or ClusterRoleBinding.
	// The type of the RBAC binding. Can be one of the following values: `RoleBinding`, or `ClusterRoleBinding`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ClustersInitParameters struct {

	// (String)
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type ClustersObservation struct {

	// (String)
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type ClustersParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	UID *string `json:"uid" tf:"uid,omitempty"`
}

type NamespacesInitParameters struct {

	// (List of String) List of images to disallow for the namespace. For example, ['nginx:latest', 'redis:latest']
	// List of images to disallow for the namespace. For example, `['nginx:latest', 'redis:latest']`
	ImagesBlacklist []*string `json:"imagesBlacklist,omitempty" tf:"images_blacklist,omitempty"`

	// (String)
	// Name of the namespace. This is the name of the Kubernetes namespace in the cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Map of String) Resource allocation for the namespace. This is a map containing the resource type and the resource value. For example, {cpu_cores: '2', memory_MiB: '2048'}
	// Resource allocation for the namespace. This is a map containing the resource type and the resource value. For example, `{cpu_cores: '2', memory_MiB: '2048'}`
	// +mapType=granular
	ResourceAllocation map[string]*string `json:"resourceAllocation,omitempty" tf:"resource_allocation,omitempty"`
}

type NamespacesObservation struct {

	// (List of String) List of images to disallow for the namespace. For example, ['nginx:latest', 'redis:latest']
	// List of images to disallow for the namespace. For example, `['nginx:latest', 'redis:latest']`
	ImagesBlacklist []*string `json:"imagesBlacklist,omitempty" tf:"images_blacklist,omitempty"`

	// (String)
	// Name of the namespace. This is the name of the Kubernetes namespace in the cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Map of String) Resource allocation for the namespace. This is a map containing the resource type and the resource value. For example, {cpu_cores: '2', memory_MiB: '2048'}
	// Resource allocation for the namespace. This is a map containing the resource type and the resource value. For example, `{cpu_cores: '2', memory_MiB: '2048'}`
	// +mapType=granular
	ResourceAllocation map[string]*string `json:"resourceAllocation,omitempty" tf:"resource_allocation,omitempty"`
}

type NamespacesParameters struct {

	// (List of String) List of images to disallow for the namespace. For example, ['nginx:latest', 'redis:latest']
	// List of images to disallow for the namespace. For example, `['nginx:latest', 'redis:latest']`
	// +kubebuilder:validation:Optional
	ImagesBlacklist []*string `json:"imagesBlacklist,omitempty" tf:"images_blacklist,omitempty"`

	// (String)
	// Name of the namespace. This is the name of the Kubernetes namespace in the cluster.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (Map of String) Resource allocation for the namespace. This is a map containing the resource type and the resource value. For example, {cpu_cores: '2', memory_MiB: '2048'}
	// Resource allocation for the namespace. This is a map containing the resource type and the resource value. For example, `{cpu_cores: '2', memory_MiB: '2048'}`
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ResourceAllocation map[string]*string `json:"resourceAllocation" tf:"resource_allocation,omitempty"`
}

type SubjectsInitParameters struct {

	// (String)
	// The name of the subject. Required if 'type' is set to 'User' or 'Group'.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// The Kubernetes namespace of the subject. Required if 'type' is set to 'ServiceAccount'.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) The type of the RBAC binding. Can be one of the following values: RoleBinding, or ClusterRoleBinding.
	// The type of the subject. Can be one of the following values: `User`, `Group`, or `ServiceAccount`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubjectsObservation struct {

	// (String)
	// The name of the subject. Required if 'type' is set to 'User' or 'Group'.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// The Kubernetes namespace of the subject. Required if 'type' is set to 'ServiceAccount'.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) The type of the RBAC binding. Can be one of the following values: RoleBinding, or ClusterRoleBinding.
	// The type of the subject. Can be one of the following values: `User`, `Group`, or `ServiceAccount`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubjectsParameters struct {

	// (String)
	// The name of the subject. Required if 'type' is set to 'User' or 'Group'.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// The Kubernetes namespace of the subject. Required if 'type' is set to 'ServiceAccount'.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) The type of the RBAC binding. Can be one of the following values: RoleBinding, or ClusterRoleBinding.
	// The type of the subject. Can be one of the following values: `User`, `Group`, or `ServiceAccount`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type WorkspaceInitParameters_2 struct {

	// (Block List, Max: 1) The backup policy for the cluster. If not specified, no backups will be taken. (see below for nested schema)
	// The backup policy for the cluster. If not specified, no backups will be taken.
	BackupPolicy []BackupPolicyInitParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// (Block List) The RBAC binding for the cluster. (see below for nested schema)
	// The RBAC binding for the cluster.
	ClusterRbacBinding []ClusterRbacBindingInitParameters `json:"clusterRbacBinding,omitempty" tf:"cluster_rbac_binding,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Clusters []ClustersInitParameters `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) The namespaces for the cluster. (see below for nested schema)
	// The namespaces for the cluster.
	Namespaces []NamespacesInitParameters `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// (Set of String)
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Block List, Max: 1) Workspace quota default limits assigned to the namespace. (see below for nested schema)
	// Workspace quota default limits assigned to the namespace.
	WorkspaceQuota []WorkspaceQuotaInitParameters `json:"workspaceQuota,omitempty" tf:"workspace_quota,omitempty"`
}

type WorkspaceObservation_2 struct {

	// (Block List, Max: 1) The backup policy for the cluster. If not specified, no backups will be taken. (see below for nested schema)
	// The backup policy for the cluster. If not specified, no backups will be taken.
	BackupPolicy []BackupPolicyObservation `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// (Block List) The RBAC binding for the cluster. (see below for nested schema)
	// The RBAC binding for the cluster.
	ClusterRbacBinding []ClusterRbacBindingObservation `json:"clusterRbacBinding,omitempty" tf:"cluster_rbac_binding,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Clusters []ClustersObservation `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) The namespaces for the cluster. (see below for nested schema)
	// The namespaces for the cluster.
	Namespaces []NamespacesObservation `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// (Set of String)
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Block List, Max: 1) Workspace quota default limits assigned to the namespace. (see below for nested schema)
	// Workspace quota default limits assigned to the namespace.
	WorkspaceQuota []WorkspaceQuotaObservation `json:"workspaceQuota,omitempty" tf:"workspace_quota,omitempty"`
}

type WorkspaceParameters_2 struct {

	// (Block List, Max: 1) The backup policy for the cluster. If not specified, no backups will be taken. (see below for nested schema)
	// The backup policy for the cluster. If not specified, no backups will be taken.
	// +kubebuilder:validation:Optional
	BackupPolicy []BackupPolicyParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// (Block List) The RBAC binding for the cluster. (see below for nested schema)
	// The RBAC binding for the cluster.
	// +kubebuilder:validation:Optional
	ClusterRbacBinding []ClusterRbacBindingParameters `json:"clusterRbacBinding,omitempty" tf:"cluster_rbac_binding,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Clusters []ClustersParameters `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) The namespaces for the cluster. (see below for nested schema)
	// The namespaces for the cluster.
	// +kubebuilder:validation:Optional
	Namespaces []NamespacesParameters `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// (Set of String)
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Block List, Max: 1) Workspace quota default limits assigned to the namespace. (see below for nested schema)
	// Workspace quota default limits assigned to the namespace.
	// +kubebuilder:validation:Optional
	WorkspaceQuota []WorkspaceQuotaParameters `json:"workspaceQuota,omitempty" tf:"workspace_quota,omitempty"`
}

type WorkspaceQuotaInitParameters struct {

	// (Number) CPU that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	// CPU that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// (Number) Memory in Mib that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	// Memory in Mib that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`
}

type WorkspaceQuotaObservation struct {

	// (Number) CPU that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	// CPU that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// (Number) Memory in Mib that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	// Memory in Mib that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`
}

type WorkspaceQuotaParameters struct {

	// (Number) CPU that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	// CPU that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	// +kubebuilder:validation:Optional
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// (Number) Memory in Mib that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	// Memory in Mib that the entire workspace is allowed to consume. The default value is 0, which imposes no limit.
	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters_2 `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider WorkspaceInitParameters_2 `json:"initProvider,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Workspace is the Schema for the Workspaces API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusters) || (has(self.initProvider) && has(self.initProvider.clusters))",message="spec.forProvider.clusters is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   WorkspaceSpec   `json:"spec"`
	Status WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Repository type metadata.
var (
	Workspace_Kind             = "Workspace"
	Workspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workspace_Kind}.String()
	Workspace_KindAPIVersion   = Workspace_Kind + "." + CRDGroupVersion.String()
	Workspace_GroupVersionKind = CRDGroupVersion.WithKind(Workspace_Kind)
)

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
