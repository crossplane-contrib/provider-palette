/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-palette/apis/spectrocloud/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Token.
func (mg *Token) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectUIDRef,
		Selector:     mg.Spec.ForProvider.ProjectUIDSelector,
		To: reference.To{
			List:    &v1alpha1.ProjectList{},
			Managed: &v1alpha1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectUID")
	}
	mg.Spec.ForProvider.ProjectUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectUIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProjectUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ProjectUIDRef,
		Selector:     mg.Spec.InitProvider.ProjectUIDSelector,
		To: reference.To{
			List:    &v1alpha1.ProjectList{},
			Managed: &v1alpha1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProjectUID")
	}
	mg.Spec.InitProvider.ProjectUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectUIDRef = rsp.ResolvedReference

	return nil
}
