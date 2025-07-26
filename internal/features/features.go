/*
 Copyright 2022 Upbound Inc
*/

package features

import "github.com/crossplane/crossplane-runtime/pkg/feature"

// Feature flags.
const (
	// Note: EnableAlphaExternalSecretStores removed as ESS functionality was removed in Crossplane v2

	// EnableBetaManagementPolicies enables beta support for
	// Management Policies. See the below design for more details.
	// https://github.com/crossplane/crossplane/pull/3531
	EnableBetaManagementPolicies feature.Flag = "EnableBetaManagementPolicies"
)
