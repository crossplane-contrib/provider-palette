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

type ProjectRoleInitParameters struct {

	// (String) Project id to be associated with the user.
	// Project id to be associated with the user.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Set of String) List of project role ids to be associated with the user.
	// List of project role ids to be associated with the user.
	// +listType=set
	RoleIds []*string `json:"roleIds,omitempty" tf:"role_ids,omitempty"`
}

type ProjectRoleObservation struct {

	// (String) Project id to be associated with the user.
	// Project id to be associated with the user.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Set of String) List of project role ids to be associated with the user.
	// List of project role ids to be associated with the user.
	// +listType=set
	RoleIds []*string `json:"roleIds,omitempty" tf:"role_ids,omitempty"`
}

type ProjectRoleParameters struct {

	// (String) Project id to be associated with the user.
	// Project id to be associated with the user.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// (Set of String) List of project role ids to be associated with the user.
	// List of project role ids to be associated with the user.
	// +kubebuilder:validation:Optional
	// +listType=set
	RoleIds []*string `json:"roleIds" tf:"role_ids,omitempty"`
}

type ResourceRoleInitParameters struct {

	// (Set of String) List of filter ids.
	// List of filter ids.
	// +listType=set
	FilterIds []*string `json:"filterIds,omitempty" tf:"filter_ids,omitempty"`

	// (Set of String) Project id's to be associated with the user.
	// Project id's to be associated with the user.
	// +listType=set
	ProjectIds []*string `json:"projectIds,omitempty" tf:"project_ids,omitempty"`

	// (Set of String) List of project role ids to be associated with the user.
	// List of resource role ids to be associated with the user.
	// +listType=set
	RoleIds []*string `json:"roleIds,omitempty" tf:"role_ids,omitempty"`
}

type ResourceRoleObservation struct {

	// (Set of String) List of filter ids.
	// List of filter ids.
	// +listType=set
	FilterIds []*string `json:"filterIds,omitempty" tf:"filter_ids,omitempty"`

	// (Set of String) Project id's to be associated with the user.
	// Project id's to be associated with the user.
	// +listType=set
	ProjectIds []*string `json:"projectIds,omitempty" tf:"project_ids,omitempty"`

	// (Set of String) List of project role ids to be associated with the user.
	// List of resource role ids to be associated with the user.
	// +listType=set
	RoleIds []*string `json:"roleIds,omitempty" tf:"role_ids,omitempty"`
}

type ResourceRoleParameters struct {

	// (Set of String) List of filter ids.
	// List of filter ids.
	// +kubebuilder:validation:Optional
	// +listType=set
	FilterIds []*string `json:"filterIds" tf:"filter_ids,omitempty"`

	// (Set of String) Project id's to be associated with the user.
	// Project id's to be associated with the user.
	// +kubebuilder:validation:Optional
	// +listType=set
	ProjectIds []*string `json:"projectIds" tf:"project_ids,omitempty"`

	// (Set of String) List of project role ids to be associated with the user.
	// List of resource role ids to be associated with the user.
	// +kubebuilder:validation:Optional
	// +listType=set
	RoleIds []*string `json:"roleIds" tf:"role_ids,omitempty"`
}

type UserInitParameters struct {

	// (String) The email of the user.
	// The email of the user.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) The first name of the user.
	// The first name of the user.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// (String) The last name of the user.
	// The last name of the user.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// (Block Set) List of project roles to be associated with the user. (see below for nested schema)
	// List of project roles to be associated with the user.
	ProjectRole []ProjectRoleInitParameters `json:"projectRole,omitempty" tf:"project_role,omitempty"`

	// (Block Set) (see below for nested schema)
	ResourceRole []ResourceRoleInitParameters `json:"resourceRole,omitempty" tf:"resource_role,omitempty"`

	// (List of String) The team id's assigned to the user.
	// The team id's assigned to the user.
	TeamIds []*string `json:"teamIds,omitempty" tf:"team_ids,omitempty"`

	// (Set of String) List of tenant role ids to be associated with the user.
	// List of tenant role ids to be associated with the user.
	// +listType=set
	TenantRole []*string `json:"tenantRole,omitempty" tf:"tenant_role,omitempty"`

	// (Block Set) List of workspace roles to be associated with the user. (see below for nested schema)
	// List of workspace roles to be associated with the user.
	WorkspaceRole []WorkspaceRoleInitParameters `json:"workspaceRole,omitempty" tf:"workspace_role,omitempty"`
}

type UserObservation struct {

	// (String) The email of the user.
	// The email of the user.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) The first name of the user.
	// The first name of the user.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The last name of the user.
	// The last name of the user.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// (Block Set) List of project roles to be associated with the user. (see below for nested schema)
	// List of project roles to be associated with the user.
	ProjectRole []ProjectRoleObservation `json:"projectRole,omitempty" tf:"project_role,omitempty"`

	// (Block Set) (see below for nested schema)
	ResourceRole []ResourceRoleObservation `json:"resourceRole,omitempty" tf:"resource_role,omitempty"`

	// (List of String) The team id's assigned to the user.
	// The team id's assigned to the user.
	TeamIds []*string `json:"teamIds,omitempty" tf:"team_ids,omitempty"`

	// (Set of String) List of tenant role ids to be associated with the user.
	// List of tenant role ids to be associated with the user.
	// +listType=set
	TenantRole []*string `json:"tenantRole,omitempty" tf:"tenant_role,omitempty"`

	// (Block Set) List of workspace roles to be associated with the user. (see below for nested schema)
	// List of workspace roles to be associated with the user.
	WorkspaceRole []WorkspaceRoleObservation `json:"workspaceRole,omitempty" tf:"workspace_role,omitempty"`
}

type UserParameters struct {

	// (String) The email of the user.
	// The email of the user.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) The first name of the user.
	// The first name of the user.
	// +kubebuilder:validation:Optional
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// (String) The last name of the user.
	// The last name of the user.
	// +kubebuilder:validation:Optional
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// (Block Set) List of project roles to be associated with the user. (see below for nested schema)
	// List of project roles to be associated with the user.
	// +kubebuilder:validation:Optional
	ProjectRole []ProjectRoleParameters `json:"projectRole,omitempty" tf:"project_role,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	ResourceRole []ResourceRoleParameters `json:"resourceRole,omitempty" tf:"resource_role,omitempty"`

	// (List of String) The team id's assigned to the user.
	// The team id's assigned to the user.
	// +kubebuilder:validation:Optional
	TeamIds []*string `json:"teamIds,omitempty" tf:"team_ids,omitempty"`

	// (Set of String) List of tenant role ids to be associated with the user.
	// List of tenant role ids to be associated with the user.
	// +kubebuilder:validation:Optional
	// +listType=set
	TenantRole []*string `json:"tenantRole,omitempty" tf:"tenant_role,omitempty"`

	// (Block Set) List of workspace roles to be associated with the user. (see below for nested schema)
	// List of workspace roles to be associated with the user.
	// +kubebuilder:validation:Optional
	WorkspaceRole []WorkspaceRoleParameters `json:"workspaceRole,omitempty" tf:"workspace_role,omitempty"`
}

type WorkspaceRoleInitParameters struct {

	// (String) Project id to be associated with the user.
	// Project id to be associated with the user.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Block Set, Min: 1) List of workspace roles to be associated with the user. (see below for nested schema)
	// List of workspace roles to be associated with the user.
	Workspace []WorkspaceRoleWorkspaceInitParameters `json:"workspace,omitempty" tf:"workspace,omitempty"`
}

type WorkspaceRoleObservation struct {

	// (String) Project id to be associated with the user.
	// Project id to be associated with the user.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Block Set, Min: 1) List of workspace roles to be associated with the user. (see below for nested schema)
	// List of workspace roles to be associated with the user.
	Workspace []WorkspaceRoleWorkspaceObservation `json:"workspace,omitempty" tf:"workspace,omitempty"`
}

type WorkspaceRoleParameters struct {

	// (String) Project id to be associated with the user.
	// Project id to be associated with the user.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// (Block Set, Min: 1) List of workspace roles to be associated with the user. (see below for nested schema)
	// List of workspace roles to be associated with the user.
	// +kubebuilder:validation:Optional
	Workspace []WorkspaceRoleWorkspaceParameters `json:"workspace" tf:"workspace,omitempty"`
}

type WorkspaceRoleWorkspaceInitParameters struct {

	// (String) The ID of this resource.
	// Workspace id to be associated with the user.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) List of project role ids to be associated with the user.
	// List of workspace role ids to be associated with the user.
	// +listType=set
	RoleIds []*string `json:"roleIds,omitempty" tf:"role_ids,omitempty"`
}

type WorkspaceRoleWorkspaceObservation struct {

	// (String) The ID of this resource.
	// Workspace id to be associated with the user.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) List of project role ids to be associated with the user.
	// List of workspace role ids to be associated with the user.
	// +listType=set
	RoleIds []*string `json:"roleIds,omitempty" tf:"role_ids,omitempty"`
}

type WorkspaceRoleWorkspaceParameters struct {

	// (String) The ID of this resource.
	// Workspace id to be associated with the user.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// (Set of String) List of project role ids to be associated with the user.
	// List of workspace role ids to be associated with the user.
	// +kubebuilder:validation:Optional
	// +listType=set
	RoleIds []*string `json:"roleIds" tf:"role_ids,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// User is the Schema for the Users API. Create and manage projects in Palette.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.firstName) || (has(self.initProvider) && has(self.initProvider.firstName))",message="spec.forProvider.firstName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lastName) || (has(self.initProvider) && has(self.initProvider.lastName))",message="spec.forProvider.lastName is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}