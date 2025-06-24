#!/bin/bash

# Check if KUBECONFIG is set
if [ -z "$KUBECONFIG" ]; then
    echo "Error: KUBECONFIG environment variable is not set"
    echo "Please set it with your Kubernetes configuration path"
    exit 1
fi

export UPBOUND_CONTEXT="local"

# Check if TF_REATTACH_PROVIDERS is set
if [ -z "$TF_REATTACH_PROVIDERS" ]; then
    echo "Error: TF_REATTACH_PROVIDERS environment variable is not set"
    echo "Please set it with your Terraform provider configuration"
    exit 1
fi

# Install Delve if not already installed
if ! command -v dlv &> /dev/null; then
    echo "Installing Delve..."
    go install github.com/go-delve/delve/cmd/dlv@latest
fi

echo "Starting provider-palette debug process..."
echo "Process ID will be shown below. Use this ID when attaching the debugger."

# One-liner to kill the process using port 8080 
# lsof -i :8080 | grep LISTEN | awk '{print $2}' | xargs kill -9 

# Run the provider with Delve in headless mode
dlv debug --headless --listen=:2346 --api-version=2 --accept-multiclient cmd/provider/main.go -- \
    --debug \
    --terraform-version=1.3.3 \
    --terraform-provider-version=0.23.5 \
    --terraform-provider-source=spectrocloud/spectrocloud 