# Provider Palette

`provider-palette` is a [Crossplane](https://crossplane.io/) provider built using [Upjet](https://github.com/crossplane/upjet) code generation tools. It exposes XRM-conformant managed resources for the [Palette API](https://docs.spectrocloud.com/api/). `provider-palette`, manages Palette resources declaratively using the [Spectro Cloud Palette Terraform provider](https://github.com/spectrocloud/terraform-provider-spectrocloud).

## Getting Started

### Prerequisites
- A kubernetes cluster with crossplane pre-installed OR install [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- Install [helm](https://helm.sh/docs/intro/install/)
- Install [kubectl](https://kubernetes.io/docs/tasks/tools/)
- Install [up](https://docs.upbound.io/reference/cli/) (optional)

### Instructions

1. If you have access to a kubernetes cluster with crossplane pre-installed, ensure that your `KUBECONFIG` is properly configured. Otherwise, deploy a kind cluster locally, then follow the [crossplane installation instructions](https://docs.crossplane.io/latest/software/install/).

2. Create a secret containing your Palette credentials:
   ```
   export PALETTE_API_KEY="" # Edit me
   export PALETTE_PROJECT_NAME=Default
   export PALETTE_HOST=api.spectrocloud.com

   cat <<EOF | kubectl apply -f -
   apiVersion: v1
   kind: Secret
   metadata:
     name: palette-creds
     namespace: crossplane-system
   type: Opaque
   stringData:
     credentials: |
       {
       "api_key": "$PALETTE_API_KEY",
       "project_name": "$PALETTE_PROJECT_NAME",
       "host": "$PALETTE_HOST"
       }
   EOF
   ```

3. Install `provider-palette` using the following command:
   ```
   up ctp provider install crossplane-contrib/provider-palette:v0.19.2
   ```

   Alternatively, you may install by creating a `Provider` resource:
   ```
   cat <<EOF | kubectl apply -f -
   apiVersion: pkg.crossplane.io/v1
   kind: Provider
   metadata:
     name: provider-palette
   spec:
     package: crossplane-contrib/provider-palette:v0.19.2
   EOF
   ```
> [!IMPORTANT]
> Ensure the image tag you use matches the [latest release](https://marketplace.upbound.io/providers/crossplane-contrib/provider-palette)

4. Create a ProviderConfig to authenticate the Provider:
   ```
   cat <<EOF | kubectl apply -f -
   apiVersion: palette.crossplane.io/v1beta1
   kind: ProviderConfig
   metadata:
     name: provider-palette-config
   spec:
     credentials:
       source: Secret
       secretRef:
         name: palette-creds
         namespace: crossplane-system
         key: credentials
   EOF
   ```

5. Refer to the [examples](./examples/), [examples-generated](./examples-generated/), and [API reference](https://doc.crds.dev/github.com/crossplane-contrib/provider-palette) to learn how to start creating Palette resources.

## Developing

Before getting started, you should read the [Upjet docs](https://github.com/crossplane/upjet/tree/main/docs). At a minimum, read about [generating a provider](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md).

### Generate/regenerate `provider-palette`
```console
# Install dependencies
go install golang.org/x/tools/cmd/goimports@latest

# Prepare repo for generation
make submodules vendor vendor.check

# Generate provider-palette. To upgrade, you must first edit TERRAFORM_PROVIDER_VERSION in the Makefile.
make generate
```

### Misc. development actions
- Run `provider-palette` in debug mode against a Kubernetes cluster:
  ```
  make run
  ```
- Build, push, and install:
  ```
  make all
  ```
- Build xpkgs:
  ```
  make -j2 build
  ```

### Debug with VSCode
1. Install CRDs
   ```console
   kubectl apply -f package/crds/
   ```
2. Edit the `Debug` configuration in `.vscode/launch.json`:
   - Ensure `env.KUBECONFIG` is correct
   - Optionally specify `env.TF_REATTACH_PROVIDERS` to hook into your dev Terraform provider
3. Execute the `Debug` launch configuration

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-palette/issues).
