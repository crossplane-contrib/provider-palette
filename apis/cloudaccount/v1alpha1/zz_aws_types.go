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

type AwsInitParameters struct {

	// (String) The Amazon Resource Name (ARN) associated with the AWS resource. This is used for identifying resources in AWS.
	// The Amazon Resource Name (ARN) associated with the AWS resource. This is used for identifying resources in AWS.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// (String) The AWS access key used to authenticate.
	// The AWS access key used to authenticate.
	AwsAccessKey *string `json:"awsAccessKey,omitempty" tf:"aws_access_key,omitempty"`

	// (String) The context of the AWS configuration. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the AWS configuration. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// us-gov' for AWS GovCloud (US) regions.
	// Default is 'aws'.
	// The AWS partition in which the cloud account is located.
	// Can be 'aws' for standard AWS regions or 'aws-us-gov' for AWS GovCloud (US) regions.
	// Default is 'aws'.
	Partition *string `json:"partition,omitempty" tf:"partition,omitempty"`

	// (Set of String) A set of ARNs for the IAM policies that should be associated with the cloud account.
	// A set of ARNs for the IAM policies that should be associated with the cloud account.
	// +listType=set
	PolicyArns []*string `json:"policyArns,omitempty" tf:"policy_arns,omitempty"`

	// (String) ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	// ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	PrivateCloudGatewayID *string `json:"privateCloudGatewayId,omitempty" tf:"private_cloud_gateway_id,omitempty"`

	// (String) The type of AWS credentials to use. Can be secret or sts.
	// The type of AWS credentials to use. Can be `secret` or `sts`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AwsObservation struct {

	// (String) The Amazon Resource Name (ARN) associated with the AWS resource. This is used for identifying resources in AWS.
	// The Amazon Resource Name (ARN) associated with the AWS resource. This is used for identifying resources in AWS.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// (String) The AWS access key used to authenticate.
	// The AWS access key used to authenticate.
	AwsAccessKey *string `json:"awsAccessKey,omitempty" tf:"aws_access_key,omitempty"`

	// (String) The context of the AWS configuration. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the AWS configuration. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// us-gov' for AWS GovCloud (US) regions.
	// Default is 'aws'.
	// The AWS partition in which the cloud account is located.
	// Can be 'aws' for standard AWS regions or 'aws-us-gov' for AWS GovCloud (US) regions.
	// Default is 'aws'.
	Partition *string `json:"partition,omitempty" tf:"partition,omitempty"`

	// (Set of String) A set of ARNs for the IAM policies that should be associated with the cloud account.
	// A set of ARNs for the IAM policies that should be associated with the cloud account.
	// +listType=set
	PolicyArns []*string `json:"policyArns,omitempty" tf:"policy_arns,omitempty"`

	// (String) ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	// ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	PrivateCloudGatewayID *string `json:"privateCloudGatewayId,omitempty" tf:"private_cloud_gateway_id,omitempty"`

	// (String) The type of AWS credentials to use. Can be secret or sts.
	// The type of AWS credentials to use. Can be `secret` or `sts`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AwsParameters struct {

	// (String) The Amazon Resource Name (ARN) associated with the AWS resource. This is used for identifying resources in AWS.
	// The Amazon Resource Name (ARN) associated with the AWS resource. This is used for identifying resources in AWS.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// (String) The AWS access key used to authenticate.
	// The AWS access key used to authenticate.
	// +kubebuilder:validation:Optional
	AwsAccessKey *string `json:"awsAccessKey,omitempty" tf:"aws_access_key,omitempty"`

	// (String, Sensitive) The AWS secret key used in conjunction with the access key for authentication.
	// The AWS secret key used in conjunction with the access key for authentication.
	// +kubebuilder:validation:Optional
	AwsSecretKeySecretRef *v1.SecretKeySelector `json:"awsSecretKeySecretRef,omitempty" tf:"-"`

	// (String) The context of the AWS configuration. Allowed values are project or tenant. Default value is project. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the AWS configuration. Allowed values are `project` or `tenant`. Default value is `project`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// account access in AWS.
	// An optional external ID that can be used for cross-account access in AWS.
	// +kubebuilder:validation:Optional
	ExternalIDSecretRef *v1.SecretKeySelector `json:"externalIdSecretRef,omitempty" tf:"-"`

	// us-gov' for AWS GovCloud (US) regions.
	// Default is 'aws'.
	// The AWS partition in which the cloud account is located.
	// Can be 'aws' for standard AWS regions or 'aws-us-gov' for AWS GovCloud (US) regions.
	// Default is 'aws'.
	// +kubebuilder:validation:Optional
	Partition *string `json:"partition,omitempty" tf:"partition,omitempty"`

	// (Set of String) A set of ARNs for the IAM policies that should be associated with the cloud account.
	// A set of ARNs for the IAM policies that should be associated with the cloud account.
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyArns []*string `json:"policyArns,omitempty" tf:"policy_arns,omitempty"`

	// (String) ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	// ID of the private cloud gateway. This is the ID of the private cloud gateway that is used to connect to the private cluster endpoint.
	// +kubebuilder:validation:Optional
	PrivateCloudGatewayID *string `json:"privateCloudGatewayId,omitempty" tf:"private_cloud_gateway_id,omitempty"`

	// (String) The type of AWS credentials to use. Can be secret or sts.
	// The type of AWS credentials to use. Can be `secret` or `sts`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// AwsSpec defines the desired state of Aws
type AwsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AwsParameters `json:"forProvider"`
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
	InitProvider AwsInitParameters `json:"initProvider,omitempty"`
}

// AwsStatus defines the observed state of Aws.
type AwsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AwsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Aws is the Schema for the Awss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Aws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSpec   `json:"spec"`
	Status            AwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwsList contains a list of Awss
type AwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Aws `json:"items"`
}

// Repository type metadata.
var (
	Aws_Kind             = "Aws"
	Aws_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Aws_Kind}.String()
	Aws_KindAPIVersion   = Aws_Kind + "." + CRDGroupVersion.String()
	Aws_GroupVersionKind = CRDGroupVersion.WithKind(Aws_Kind)
)

func init() {
	SchemeBuilder.Register(&Aws{}, &AwsList{})
}
