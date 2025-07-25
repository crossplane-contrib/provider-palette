/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-palette/apis/backup/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-palette/apis/cluster/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.BackupPolicy); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupPolicy[i3].BackupLocationID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.BackupPolicy[i3].BackupLocationIDRef,
			Selector:     mg.Spec.ForProvider.BackupPolicy[i3].BackupLocationIDSelector,
			To: reference.To{
				List:    &v1alpha1.StorageLocationList{},
				Managed: &v1alpha1.StorageLocation{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BackupPolicy[i3].BackupLocationID")
		}
		mg.Spec.ForProvider.BackupPolicy[i3].BackupLocationID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.BackupPolicy[i3].BackupLocationIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterGroupUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterGroupUIDRef,
		Selector:     mg.Spec.ForProvider.ClusterGroupUIDSelector,
		To: reference.To{
			List:    &v1alpha11.GroupList{},
			Managed: &v1alpha11.Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterGroupUID")
	}
	mg.Spec.ForProvider.ClusterGroupUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterGroupUIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ClusterProfile); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterProfile[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterProfile[i3].IDRef,
			Selector:     mg.Spec.ForProvider.ClusterProfile[i3].IDSelector,
			To: reference.To{
				List:    &v1alpha11.ProfileList{},
				Managed: &v1alpha11.Profile{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ClusterProfile[i3].ID")
		}
		mg.Spec.ForProvider.ClusterProfile[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ClusterProfile[i3].IDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.BackupPolicy); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BackupPolicy[i3].BackupLocationID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.BackupPolicy[i3].BackupLocationIDRef,
			Selector:     mg.Spec.InitProvider.BackupPolicy[i3].BackupLocationIDSelector,
			To: reference.To{
				List:    &v1alpha1.StorageLocationList{},
				Managed: &v1alpha1.StorageLocation{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.BackupPolicy[i3].BackupLocationID")
		}
		mg.Spec.InitProvider.BackupPolicy[i3].BackupLocationID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.BackupPolicy[i3].BackupLocationIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterGroupUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterGroupUIDRef,
		Selector:     mg.Spec.InitProvider.ClusterGroupUIDSelector,
		To: reference.To{
			List:    &v1alpha11.GroupList{},
			Managed: &v1alpha11.Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterGroupUID")
	}
	mg.Spec.InitProvider.ClusterGroupUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterGroupUIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.ClusterProfile); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterProfile[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterProfile[i3].IDRef,
			Selector:     mg.Spec.InitProvider.ClusterProfile[i3].IDSelector,
			To: reference.To{
				List:    &v1alpha11.ProfileList{},
				Managed: &v1alpha11.Profile{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ClusterProfile[i3].ID")
		}
		mg.Spec.InitProvider.ClusterProfile[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ClusterProfile[i3].IDRef = rsp.ResolvedReference

	}

	return nil
}
