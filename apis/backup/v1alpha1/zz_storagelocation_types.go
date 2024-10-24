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

type S3InitParameters struct {

	// (String) The access key for S3 authentication, required if 'credential_type' is set to 'secret'.
	// The access key for S3 authentication, required if 'credential_type' is set to 'secret'.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// (String) The Amazon Resource Name (ARN) of the IAM role to assume for accessing S3 when using 'sts' credentials.
	// The Amazon Resource Name (ARN) of the IAM role to assume for accessing S3 when using 'sts' credentials.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// based credentials.
	// The type of credentials used to access the S3 storage. Supported values are 'secret' for static credentials and 'sts' for temporary, token-based credentials.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// account access to the S3 storage when using 'sts' credentials.
	// An external ID used for cross-account access to the S3 storage when using 'sts' credentials.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// style URL for accessing S3.
	// A boolean flag indicating whether to enforce the path-style URL for accessing S3.
	S3ForcePathStyle *bool `json:"s3ForcePathStyle,omitempty" tf:"s3_force_path_style,omitempty"`

	// (String) The S3 URL endpoint.
	// The S3 URL endpoint.
	S3URL *string `json:"s3Url,omitempty" tf:"s3_url,omitempty"`

	// (String) The secret key for S3 authentication, required if 'credential_type' is set to 'secret'.
	// The secret key for S3 authentication, required if 'credential_type' is set to 'secret'.
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key,omitempty"`
}

type S3Observation struct {

	// (String) The access key for S3 authentication, required if 'credential_type' is set to 'secret'.
	// The access key for S3 authentication, required if 'credential_type' is set to 'secret'.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// (String) The Amazon Resource Name (ARN) of the IAM role to assume for accessing S3 when using 'sts' credentials.
	// The Amazon Resource Name (ARN) of the IAM role to assume for accessing S3 when using 'sts' credentials.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// based credentials.
	// The type of credentials used to access the S3 storage. Supported values are 'secret' for static credentials and 'sts' for temporary, token-based credentials.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// account access to the S3 storage when using 'sts' credentials.
	// An external ID used for cross-account access to the S3 storage when using 'sts' credentials.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// style URL for accessing S3.
	// A boolean flag indicating whether to enforce the path-style URL for accessing S3.
	S3ForcePathStyle *bool `json:"s3ForcePathStyle,omitempty" tf:"s3_force_path_style,omitempty"`

	// (String) The S3 URL endpoint.
	// The S3 URL endpoint.
	S3URL *string `json:"s3Url,omitempty" tf:"s3_url,omitempty"`

	// (String) The secret key for S3 authentication, required if 'credential_type' is set to 'secret'.
	// The secret key for S3 authentication, required if 'credential_type' is set to 'secret'.
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key,omitempty"`
}

type S3Parameters struct {

	// (String) The access key for S3 authentication, required if 'credential_type' is set to 'secret'.
	// The access key for S3 authentication, required if 'credential_type' is set to 'secret'.
	// +kubebuilder:validation:Optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// (String) The Amazon Resource Name (ARN) of the IAM role to assume for accessing S3 when using 'sts' credentials.
	// The Amazon Resource Name (ARN) of the IAM role to assume for accessing S3 when using 'sts' credentials.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// based credentials.
	// The type of credentials used to access the S3 storage. Supported values are 'secret' for static credentials and 'sts' for temporary, token-based credentials.
	// +kubebuilder:validation:Optional
	CredentialType *string `json:"credentialType" tf:"credential_type,omitempty"`

	// account access to the S3 storage when using 'sts' credentials.
	// An external ID used for cross-account access to the S3 storage when using 'sts' credentials.
	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// style URL for accessing S3.
	// A boolean flag indicating whether to enforce the path-style URL for accessing S3.
	// +kubebuilder:validation:Optional
	S3ForcePathStyle *bool `json:"s3ForcePathStyle,omitempty" tf:"s3_force_path_style,omitempty"`

	// (String) The S3 URL endpoint.
	// The S3 URL endpoint.
	// +kubebuilder:validation:Optional
	S3URL *string `json:"s3Url,omitempty" tf:"s3_url,omitempty"`

	// (String) The secret key for S3 authentication, required if 'credential_type' is set to 'secret'.
	// The secret key for S3 authentication, required if 'credential_type' is set to 'secret'.
	// +kubebuilder:validation:Optional
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key,omitempty"`
}

type StorageLocationInitParameters struct {

	// compatible storage services.
	// The name of the storage bucket where backups are stored. This is relevant for S3 or S3-compatible storage services.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String) An optional CA certificate used for SSL connections to ensure secure communication with the storage provider.
	// An optional CA certificate used for SSL connections to ensure secure communication with the storage provider.
	CACert *string `json:"caCert,omitempty" tf:"ca_cert,omitempty"`

	// (String) The context of the backup storage location. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the backup storage location. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (Boolean) Specifies if this backup storage location should be used as the default location for storing backups.
	// Specifies if this backup storage location should be used as the default location for storing backups.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// (String) The region where the backup storage is located, typically corresponding to the region of the cloud provider.
	// The region where the backup storage is located, typically corresponding to the region of the cloud provider.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// specific settings for configuring the backup storage location. (see below for nested schema)
	// S3-specific settings for configuring the backup storage location.
	S3 []S3InitParameters `json:"s3,omitempty" tf:"s3,omitempty"`
}

type StorageLocationObservation struct {

	// compatible storage services.
	// The name of the storage bucket where backups are stored. This is relevant for S3 or S3-compatible storage services.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String) An optional CA certificate used for SSL connections to ensure secure communication with the storage provider.
	// An optional CA certificate used for SSL connections to ensure secure communication with the storage provider.
	CACert *string `json:"caCert,omitempty" tf:"ca_cert,omitempty"`

	// (String) The context of the backup storage location. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the backup storage location. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Specifies if this backup storage location should be used as the default location for storing backups.
	// Specifies if this backup storage location should be used as the default location for storing backups.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// (String) The region where the backup storage is located, typically corresponding to the region of the cloud provider.
	// The region where the backup storage is located, typically corresponding to the region of the cloud provider.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// specific settings for configuring the backup storage location. (see below for nested schema)
	// S3-specific settings for configuring the backup storage location.
	S3 []S3Observation `json:"s3,omitempty" tf:"s3,omitempty"`
}

type StorageLocationParameters struct {

	// compatible storage services.
	// The name of the storage bucket where backups are stored. This is relevant for S3 or S3-compatible storage services.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String) An optional CA certificate used for SSL connections to ensure secure communication with the storage provider.
	// An optional CA certificate used for SSL connections to ensure secure communication with the storage provider.
	// +kubebuilder:validation:Optional
	CACert *string `json:"caCert,omitempty" tf:"ca_cert,omitempty"`

	// (String) The context of the backup storage location. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the backup storage location. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (Boolean) Specifies if this backup storage location should be used as the default location for storing backups.
	// Specifies if this backup storage location should be used as the default location for storing backups.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// (String) The region where the backup storage is located, typically corresponding to the region of the cloud provider.
	// The region where the backup storage is located, typically corresponding to the region of the cloud provider.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// specific settings for configuring the backup storage location. (see below for nested schema)
	// S3-specific settings for configuring the backup storage location.
	// +kubebuilder:validation:Optional
	S3 []S3Parameters `json:"s3,omitempty" tf:"s3,omitempty"`
}

// StorageLocationSpec defines the desired state of StorageLocation
type StorageLocationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageLocationParameters `json:"forProvider"`
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
	InitProvider StorageLocationInitParameters `json:"initProvider,omitempty"`
}

// StorageLocationStatus defines the observed state of StorageLocation.
type StorageLocationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageLocationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StorageLocation is the Schema for the StorageLocations API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type StorageLocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucketName) || (has(self.initProvider) && has(self.initProvider.bucketName))",message="spec.forProvider.bucketName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isDefault) || (has(self.initProvider) && has(self.initProvider.isDefault))",message="spec.forProvider.isDefault is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.s3) || (has(self.initProvider) && has(self.initProvider.s3))",message="spec.forProvider.s3 is a required parameter"
	Spec   StorageLocationSpec   `json:"spec"`
	Status StorageLocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageLocationList contains a list of StorageLocations
type StorageLocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageLocation `json:"items"`
}

// Repository type metadata.
var (
	StorageLocation_Kind             = "StorageLocation"
	StorageLocation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageLocation_Kind}.String()
	StorageLocation_KindAPIVersion   = StorageLocation_Kind + "." + CRDGroupVersion.String()
	StorageLocation_GroupVersionKind = CRDGroupVersion.WithKind(StorageLocation_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageLocation{}, &StorageLocationList{})
}
