// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	deployment "github.com/crossplane-contrib/provider-palette/internal/controller/addon/deployment"
	profile "github.com/crossplane-contrib/provider-palette/internal/controller/application/profile"
	storagelocation "github.com/crossplane-contrib/provider-palette/internal/controller/backup/storagelocation"
	aws "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/aws"
	azure "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/azure"
	custom "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/custom"
	gcp "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/gcp"
	maas "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/maas"
	openstack "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/openstack"
	tencent "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/tencent"
	vsphere "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/vsphere"
	aks "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/aks"
	awscluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/aws"
	azurecluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/azure"
	customcloud "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/customcloud"
	edgenative "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/edgenative"
	edgevsphere "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/edgevsphere"
	eks "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/eks"
	gcpcluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/gcp"
	gke "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/gke"
	group "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/group"
	libvirt "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/libvirt"
	maascluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/maas"
	openstackcluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/openstack"
	profilecluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/profile"
	tke "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/tke"
	vspherecluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/vsphere"
	ippool "github.com/crossplane-contrib/provider-palette/internal/controller/privatecloudgateway/ippool"
	providerconfig "github.com/crossplane-contrib/provider-palette/internal/controller/providerconfig"
	helm "github.com/crossplane-contrib/provider-palette/internal/controller/registry/helm"
	oci "github.com/crossplane-contrib/provider-palette/internal/controller/registry/oci"
	alert "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/alert"
	appliance "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/appliance"
	application "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/application"
	datavolume "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/datavolume"
	filter "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/filter"
	macro "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/macro"
	macros "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/macros"
	project "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/project"
	team "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/team"
	workspace "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/workspace"
	cluster "github.com/crossplane-contrib/provider-palette/internal/controller/virtual/cluster"
	machine "github.com/crossplane-contrib/provider-palette/internal/controller/virtual/machine"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		deployment.Setup,
		profile.Setup,
		storagelocation.Setup,
		aws.Setup,
		azure.Setup,
		custom.Setup,
		gcp.Setup,
		maas.Setup,
		openstack.Setup,
		tencent.Setup,
		vsphere.Setup,
		aks.Setup,
		awscluster.Setup,
		azurecluster.Setup,
		customcloud.Setup,
		edgenative.Setup,
		edgevsphere.Setup,
		eks.Setup,
		gcpcluster.Setup,
		gke.Setup,
		group.Setup,
		libvirt.Setup,
		maascluster.Setup,
		openstackcluster.Setup,
		profilecluster.Setup,
		tke.Setup,
		vspherecluster.Setup,
		ippool.Setup,
		providerconfig.Setup,
		helm.Setup,
		oci.Setup,
		alert.Setup,
		appliance.Setup,
		application.Setup,
		datavolume.Setup,
		filter.Setup,
		macro.Setup,
		macros.Setup,
		project.Setup,
		team.Setup,
		workspace.Setup,
		cluster.Setup,
		machine.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
