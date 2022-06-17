/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	storagelocation "github.com/crossplane-contrib/provider-jet-palette/internal/controller/backup/storagelocation"
	aws "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cloudaccount/aws"
	azure "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cloudaccount/azure"
	gcp "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cloudaccount/gcp"
	maas "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cloudaccount/maas"
	openstack "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cloudaccount/openstack"
	tencent "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cloudaccount/tencent"
	aks "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/aks"
	awscluster "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/aws"
	azurecluster "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/azure"
	edge "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/edge"
	edgevsphere "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/edgevsphere"
	eks "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/eks"
	gcpcluster "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/gcp"
	libvirt "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/libvirt"
	maascluster "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/maas"
	openstackcluster "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/openstack"
	profile "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/profile"
	tke "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/tke"
	vsphere "github.com/crossplane-contrib/provider-jet-palette/internal/controller/cluster/vsphere"
	ippool "github.com/crossplane-contrib/provider-jet-palette/internal/controller/privatecloudgateway/ippool"
	providerconfig "github.com/crossplane-contrib/provider-jet-palette/internal/controller/providerconfig"
	helm "github.com/crossplane-contrib/provider-jet-palette/internal/controller/registry/helm"
	oci "github.com/crossplane-contrib/provider-jet-palette/internal/controller/registry/oci"
	appliance "github.com/crossplane-contrib/provider-jet-palette/internal/controller/spectrocloud/appliance"
	macro "github.com/crossplane-contrib/provider-jet-palette/internal/controller/spectrocloud/macro"
	project "github.com/crossplane-contrib/provider-jet-palette/internal/controller/spectrocloud/project"
	team "github.com/crossplane-contrib/provider-jet-palette/internal/controller/spectrocloud/team"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		storagelocation.Setup,
		aws.Setup,
		azure.Setup,
		gcp.Setup,
		maas.Setup,
		openstack.Setup,
		tencent.Setup,
		aks.Setup,
		awscluster.Setup,
		azurecluster.Setup,
		edge.Setup,
		edgevsphere.Setup,
		eks.Setup,
		gcpcluster.Setup,
		libvirt.Setup,
		maascluster.Setup,
		openstackcluster.Setup,
		profile.Setup,
		tke.Setup,
		vsphere.Setup,
		ippool.Setup,
		providerconfig.Setup,
		helm.Setup,
		oci.Setup,
		appliance.Setup,
		macro.Setup,
		project.Setup,
		team.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
