//go:build generate

/*
Copyright 2021 Upbound Inc.
*/

// Remove existing CRDs before applying new ones.
//go:generate rm -f ../pkg/crds/*.yaml

// Generate deepcopy methodsets and CRD manifests
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./... crd:allowDangerousTypes=true,crdVersions=v1 output:crd:artifacts:config=../pkg/crds

// RBAC rules are manually created in rbac.yaml in the root directory
// The controller-gen RBAC generator doesn't work well with upjet-generated types

package apis
