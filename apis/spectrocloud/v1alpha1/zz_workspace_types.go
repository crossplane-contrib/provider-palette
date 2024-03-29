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

type BackupPolicyObservation struct {
}

type BackupPolicyParameters struct {

	// The ID of the backup location to use for the backup.
	// +kubebuilder:validation:Required
	BackupLocationID *string `json:"backupLocationId" tf:"backup_location_id,omitempty"`

	// The list of cluster UIDs to include in the backup. If `include_all_clusters` is set to `true`, then all clusters will be included.
	// +kubebuilder:validation:Optional
	ClusterUids []*string `json:"clusterUids,omitempty" tf:"cluster_uids,omitempty"`

	// The number of hours after which the backup will be deleted. For example, if the expiry is set to 24, the backup will be deleted after 24 hours.
	// +kubebuilder:validation:Required
	ExpiryInHour *float64 `json:"expiryInHour" tf:"expiry_in_hour,omitempty"`

	// Whether to include all clusters in the backup. If set to false, only the clusters specified in `cluster_uids` will be included.
	// +kubebuilder:validation:Optional
	IncludeAllClusters *bool `json:"includeAllClusters,omitempty" tf:"include_all_clusters,omitempty"`

	// Whether to include the cluster resources in the backup. If set to false, only the cluster configuration and disks will be backed up.
	// +kubebuilder:validation:Optional
	IncludeClusterResources *bool `json:"includeClusterResources,omitempty" tf:"include_cluster_resources,omitempty"`

	// Whether to include the disks in the backup. If set to false, only the cluster configuration will be backed up.
	// +kubebuilder:validation:Optional
	IncludeDisks *bool `json:"includeDisks,omitempty" tf:"include_disks,omitempty"`

	// The list of Kubernetes namespaces to include in the backup. If not specified, all namespaces will be included.
	// +kubebuilder:validation:Optional
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// Prefix for the backup name. The backup name will be of the format <prefix>-<cluster-name>-<timestamp>.
	// +kubebuilder:validation:Required
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// The schedule for the backup. The schedule is specified in cron format. For example, to run the backup every day at 1:00 AM, the schedule should be set to `0 1 * * *`.
	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type ClusterRbacBindingObservation struct {
}

type ClusterRbacBindingParameters struct {

	// The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The role of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// +kubebuilder:validation:Optional
	Role map[string]*string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Subjects []SubjectsParameters `json:"subjects,omitempty" tf:"subjects,omitempty"`

	// The type of the RBAC binding. Can be one of the following values: `RoleBinding`, or `ClusterRoleBinding`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ClustersObservation struct {
}

type ClustersParameters struct {

	// +kubebuilder:validation:Required
	UID *string `json:"uid" tf:"uid,omitempty"`
}

type NamespacesObservation struct {
}

type NamespacesParameters struct {

	// List of images to disallow for the namespace. For example, `['nginx:latest', 'redis:latest']`
	// +kubebuilder:validation:Optional
	ImagesBlacklist []*string `json:"imagesBlacklist,omitempty" tf:"images_blacklist,omitempty"`

	// Name of the namespace. This is the name of the Kubernetes namespace in the cluster.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Resource allocation for the namespace. This is a map containing the resource type and the resource value. For example, `{cpu_cores: '2', memory_MiB: '2048'}`
	// +kubebuilder:validation:Required
	ResourceAllocation map[string]*string `json:"resourceAllocation" tf:"resource_allocation,omitempty"`
}

type SubjectsObservation struct {
}

type SubjectsParameters struct {

	// The name of the subject. Required if 'type' is set to 'User' or 'Group'.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The Kubernetes namespace of the subject. Required if 'type' is set to 'ServiceAccount'.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The type of the subject. Can be one of the following values: `User`, `Group`, or `ServiceAccount`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type WorkspaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WorkspaceParameters struct {

	// +kubebuilder:validation:Optional
	BackupPolicy []BackupPolicyParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterRbacBinding []ClusterRbacBindingParameters `json:"clusterRbacBinding,omitempty" tf:"cluster_rbac_binding,omitempty"`

	// +kubebuilder:validation:Required
	Clusters []ClustersParameters `json:"clusters" tf:"clusters,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Namespaces []NamespacesParameters `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workspace is the Schema for the Workspaces API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec"`
	Status            WorkspaceStatus `json:"status,omitempty"`
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
