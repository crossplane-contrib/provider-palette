# ====================================================================================
# Setup Project

PROJECT_NAME := provider-palette
PROJECT_REPO := github.com/crossplane-contrib/$(PROJECT_NAME)

export TERRAFORM_VERSION := 1.3.3
export TERRAFORM_PROVIDER_SOURCE := spectrocloud/spectrocloud
export TERRAFORM_PROVIDER_REPO := https://github.com/spectrocloud/terraform-provider-spectrocloud
export TERRAFORM_PROVIDER_VERSION := 0.23.8
export TERRAFORM_PROVIDER_DOWNLOAD_NAME := terraform-provider-spectrocloud
export TERRAFORM_NATIVE_PROVIDER_BINARY := terraform-provider-spectrocloud_$(TERRAFORM_PROVIDER_VERSION)
export TERRAFORM_DOCS_PATH := docs/resources

PLATFORMS ?= linux_amd64 linux_arm64

# Build registry for Docker images
BUILD_REGISTRY ?= build-$(shell echo $(HOSTNAME)-$(ROOT_DIR) | shasum | cut -c1-8)

# -include will silently skip missing files, which allows us
# to load those files with a target in the Makefile. If only
# "include" was used, the make command would fail and refuse
# to run a target until the include commands succeeded.
-include build/makelib/common.mk

# ====================================================================================
# Setup Output

-include build/makelib/output.mk

# ====================================================================================
# Setup Go

# Set a sane default so that the nprocs calculation below is less noisy on the initial
# loading of this file
NPROCS ?= 1

# each of our test suites starts a kube-apiserver and running many test suites in
# parallel can lead to high CPU utilization. by default we reduce the parallelism
# to half the number of CPU cores.
GO_TEST_PARALLEL := $(shell echo $$(( $(NPROCS) / 2 )))

GO_REQUIRED_VERSION ?= 1.21
GOLANGCILINT_VERSION ?= 1.64.6
GO_STATIC_PACKAGES = $(GO_PROJECT)/cmd/provider $(GO_PROJECT)/cmd/generator
GO_LDFLAGS += -X $(GO_PROJECT)/internal/version.Version=$(VERSION)
GO_SUBDIRS += cmd internal apis
-include build/makelib/golang.mk

# ====================================================================================
# Setup Terraform for fetching provider schema
TERRAFORM := $(TOOLS_HOST_DIR)/terraform-$(TERRAFORM_VERSION)
TERRAFORM_WORKDIR := $(WORK_DIR)/terraform
TERRAFORM_PROVIDER_SCHEMA := config/schema.json

$(TERRAFORM):
	@$(INFO) installing terraform $(HOSTOS)-$(HOSTARCH)
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-terraform
	@curl -fsSL https://releases.hashicorp.com/terraform/$(TERRAFORM_VERSION)/terraform_$(TERRAFORM_VERSION)_$(SAFEHOST_PLATFORM).zip -o $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip
	@unzip $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip -d $(TOOLS_HOST_DIR)/tmp-terraform
	@mv $(TOOLS_HOST_DIR)/tmp-terraform/terraform $(TERRAFORM)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-terraform
	@$(OK) installing terraform $(HOSTOS)-$(HOSTARCH)

$(TERRAFORM_PROVIDER_SCHEMA): $(TERRAFORM)
	@$(INFO) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)
	@rm -rf $(TERRAFORM_WORKDIR)
	@mkdir -p $(TERRAFORM_WORKDIR)
	@echo '{"terraform":[{"required_providers":[{"provider":{"source":"'"$(TERRAFORM_PROVIDER_SOURCE)"'","version":"'"$(TERRAFORM_PROVIDER_VERSION)"'"}}],"required_version":"'"$(TERRAFORM_VERSION)"'"}]}' > $(TERRAFORM_WORKDIR)/main.tf.json
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) init > $(TERRAFORM_WORKDIR)/terraform-logs.txt 2>&1
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) providers schema -json=true > $(TERRAFORM_PROVIDER_SCHEMA) 2>> $(TERRAFORM_WORKDIR)/terraform-logs.txt
	@$(OK) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)

pull-docs:
	@if [ ! -d "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" ]; then \
  		mkdir -p "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" && \
		git clone -c advice.detachedHead=false --depth 1 --filter=blob:none --branch "v$(TERRAFORM_PROVIDER_VERSION)" --sparse "$(TERRAFORM_PROVIDER_REPO)" "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)"; \
	fi
	@git -C "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" sparse-checkout set "$(TERRAFORM_DOCS_PATH)"

generate.init: $(TERRAFORM_PROVIDER_SCHEMA) pull-docs

.PHONY: $(TERRAFORM_PROVIDER_SCHEMA) pull-docs

# ====================================================================================
# Provider Build Targets

# Build the provider package using Crossplane CLI
provider.build: go.build
	@$(INFO) Building provider package $(PROJECT_NAME)-$(VERSION).xpkg
	@mkdir -p $(OUTPUT_DIR)/xpkg
	@crossplane xpkg build \
		--embed-runtime-image $(BUILD_REGISTRY)/$(PROJECT_NAME)-$(ARCH) \
		--package-root pkg \
		--package-file $(OUTPUT_DIR)/xpkg/$(PROJECT_NAME)-$(VERSION).xpkg || $(FAIL)
	@$(OK) Built provider package $(PROJECT_NAME)-$(VERSION).xpkg

# Build the provider image
provider.image: go.build
	@$(INFO) Building provider image $(PROJECT_NAME)-$(ARCH)
	@mkdir -p cluster/images/$(PROJECT_NAME)/bin/linux_$(ARCH)
	@cp $(GO_OUT_DIR)/provider cluster/images/$(PROJECT_NAME)/bin/linux_$(ARCH)/
	@mkdir -p cluster/images/$(PROJECT_NAME)/pkg
	@cp pkg/package.yaml cluster/images/$(PROJECT_NAME)/pkg/
	@docker build \
		--build-arg TARGETOS=linux \
		--build-arg TARGETARCH=$(ARCH) \
		--build-arg TERRAFORM_VERSION=$(TERRAFORM_VERSION) \
		--build-arg TERRAFORM_PROVIDER_SOURCE=$(TERRAFORM_PROVIDER_SOURCE) \
		--build-arg TERRAFORM_PROVIDER_VERSION=$(TERRAFORM_PROVIDER_VERSION) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_NAME=$(TERRAFORM_PROVIDER_DOWNLOAD_NAME) \
		--build-arg TERRAFORM_NATIVE_PROVIDER_BINARY=$(TERRAFORM_NATIVE_PROVIDER_BINARY) \
		-t $(BUILD_REGISTRY)/$(PROJECT_NAME)-$(ARCH) \
		-f cluster/images/$(PROJECT_NAME)/Dockerfile \
		cluster/images/$(PROJECT_NAME) || $(FAIL)
	@$(OK) Built provider image $(PROJECT_NAME)-$(ARCH)

# Build everything needed for the provider
provider: generate provider.image provider.build

# ====================================================================================
# Targets

# NOTE: the build submodule currently overrides XDG_CACHE_HOME in order to
# force the Helm 3 to use the .work/helm directory. This causes Go on Linux
# machines to use that directory as the build cache as well. We should adjust
# this behavior in the build submodule because it is also causing Linux users
# to duplicate their build cache, but for now we just make it easier to identify
# its location in CI so that we cache between builds.
go.cachedir:
	@go env GOCACHE

go.mod.cachedir:
	@go env GOMODCACHE

# Generate a coverage report for cobertura applying exclusions on
# - generated file
cobertura:
	@cat $(GO_TEST_OUTPUT)/coverage.txt | \
		grep -v zz_ | \
		$(GOCOVER_COBERTURA) > $(GO_TEST_OUTPUT)/cobertura-coverage.xml

# Update the submodules, such as the common build scripts.
submodules:
	@git submodule sync
	@git submodule update --init --recursive

# This is for running out-of-cluster locally, and is for convenience. Running
# this make target will print out the command which was used. For more control,
# try running the binary directly with different arguments.
run: go.build
	@$(INFO) Running Crossplane locally out-of-cluster . . .
	@# To see other arguments that can be provided, run the command with --help instead
	UPBOUND_CONTEXT="local" $(GO_OUT_DIR)/provider --debug --enable-management-policies

# ===================
# Integration Testing

test-integration: init-envtest ## Run integration tests
	@mkdir -p $(GO_TEST_OUTPUT) || $(FAIL)
	KUBEBUILDER_ASSETS=${KUBEBUILDER_ASSETS} go test -v -timeout 10m \
		-covermode=atomic -coverpkg=./... -coverprofile=$(GO_TEST_OUTPUT)/integration.out ./tests/...

init-envtest: setup-envtest
	$(TOOLS_HOST_DIR)/setup-envtest use --bin-dir $(TOOLS_HOST_DIR) $(ENVTEST_VERSION)
KUBEBUILDER_ASSETS = $(shell $(TOOLS_HOST_DIR)/setup-envtest use -p path --bin-dir $(TOOLS_HOST_DIR) $(ENVTEST_VERSION))

setup-envtest:
ifeq ("$(wildcard $(TOOLS_HOST_DIR)/setup-envtest)", "")
	go get sigs.k8s.io/controller-runtime/tools/setup-envtest
	GOBIN=$(TOOLS_HOST_DIR) go install sigs.k8s.io/controller-runtime/tools/setup-envtest
endif
SETUP_ENVTEST=$(TOOLS_HOST_DIR)/setup-envtest

# ====================================================================================
# Fallthrough

# run `make help` to see the targets and options

# We want submodules to be set up the first time `make` is run.
# We manage the build/ folder and its Makefiles as a submodule.
# The first time `make` is run, the includes of build/*.mk files will
# all fail, and this target will be run. The next time, the default as defined
# by the includes will be run instead.
fallthrough: submodules
	@echo Initial setup complete. Running make again . . .
	@make

.PHONY: cobertura submodules fallthrough run crds.clean provider provider.build provider.image

# ====================================================================================
# Special Targets

define CROSSPLANE_MAKE_HELP
Crossplane Targets:
    cobertura             Generate a coverage report for cobertura applying exclusions on generated files.
    submodules            Update the submodules, such as the common build scripts.
    run                   Run crossplane locally, out-of-cluster. Useful for development.
    provider              Build the complete provider package (image + xpkg).
    provider.build        Build the provider package (.xpkg file).
    provider.image        Build the provider Docker image.

endef
# The reason CROSSPLANE_MAKE_HELP is used instead of CROSSPLANE_HELP is because the crossplane
# binary will try to use CROSSPLANE_HELP if it is set, and this is for something different.
export CROSSPLANE_MAKE_HELP

crossplane.help:
	@echo "$$CROSSPLANE_MAKE_HELP"

help-special: crossplane.help

.PHONY: crossplane.help help-special
