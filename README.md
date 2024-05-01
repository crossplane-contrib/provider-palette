# Provider Palette

`provider-palette` is a [Crossplane](https://crossplane.io/) provider that is
built using [Upjet](https://github.com/crossplane/upjet) code generation tools and
exposes XRM-conformant managed resources for the Palette API. The provider helps
you create resources declaratively using the Spectro Cloud Palette API.

For more information on the Spectro Cloud Palette API, visit:
https://docs.spectrocloud.com/api/ 

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest
release](https://marketplace.upbound.io/providers/crossplane-contrib/provider-palette):
```
up ctp provider install crossplane-contrib/provider-palette:v0.19.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-palette
spec:
  package: crossplane-contrib/provider-palette:v0.19.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig
with debug enabled.

You can see the API reference
[here](https://doc.crds.dev/github.com/crossplane-contrib/provider-palette).

## Developing

Before getting started, you should read the [Upjet docs](https://github.com/crossplane/upjet/tree/main/docs). At a minimum, read about [generating a provider](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md).

Generate/regenerate the provider:
```console
go install golang.org/x/tools/cmd/goimports@latest

make submodules

# Generate provider-palette. To upgrade, you must first edit TERRAFORM_PROVIDER_VERSION in the Makefile.
make generate
```

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binaries:

```console
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
