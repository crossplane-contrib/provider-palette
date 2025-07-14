#!/bin/bash

# Help function
show_help() {
    echo "Usage: $0 [debug|help]"
    echo ""
    echo "Arguments:"
    echo "  debug      Run with TF_REATTACH_PROVIDERS for Terraform debugging (port 2346)"
    echo "  help       Show this help message"
    echo "  (none)     Run the Crossplane provider with verbose logging"
    echo ""
    echo "Environment variables:"
    echo "  KUBECONFIG             Required: Path to Kubernetes configuration"
    echo "  TF_REATTACH_PROVIDERS  Required for debug mode: Terraform provider configuration"
    echo ""
    echo "Examples:"
    echo "  $0           # Run Crossplane provider with verbose logging"
    echo "  $0 debug     # Run with Terraform provider debugging"
    echo ""
    echo "Debug mode is for debugging Terraform provider interactions"
}

# Check for help
if [ "$1" = "help" ] || [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    show_help
    exit 0
fi

# Check if KUBECONFIG is set
if [ -z "$KUBECONFIG" ]; then
    echo "Error: KUBECONFIG environment variable is not set"
    echo "Please set it with your Kubernetes configuration path"
    echo ""
    show_help
    exit 1
fi

export UPBOUND_CONTEXT="local"

# Provider configuration
TERRAFORM_VERSION="1.3.3"
TERRAFORM_PROVIDER_SOURCE="spectrocloud/spectrocloud"
TERRAFORM_PROVIDER_VERSION="0.23.7"

# Check if debug mode is requested
DEBUG_MODE=false
if [ "$1" = "debug" ]; then
    DEBUG_MODE=true
fi

if [ "$DEBUG_MODE" = true ]; then
    # Check if TF_REATTACH_PROVIDERS is set for debug mode
    if [ -z "$TF_REATTACH_PROVIDERS" ]; then
        echo "Error: TF_REATTACH_PROVIDERS environment variable is not set"
        echo "This is required for Terraform provider debugging"
        echo ""
        show_help
        exit 1
    fi

    # Install Delve if not already installed
    if ! command -v dlv &> /dev/null; then
        echo "Installing Delve..."
        go install github.com/go-delve/delve/cmd/dlv@latest
    fi

    echo "Starting provider-palette in TERRAFORM DEBUG mode..."
    echo "Using TF_REATTACH_PROVIDERS for Terraform provider debugging"
    echo "Debug server will be available at localhost:2346"
    echo "You can attach your debugger to this port"
    
    # One-liner to kill the process using port 2346 if needed
    # lsof -i :2346 | grep LISTEN | awk '{print $2}' | xargs kill -9 

    # Run the provider with Delve in headless mode for Terraform debugging
    dlv debug --headless --listen=:2346 --api-version=2 --accept-multiclient cmd/provider/main.go -- \
        --debug \
        --enable-management-policies \
        --terraform-version="$TERRAFORM_VERSION" \
        --terraform-provider-source="$TERRAFORM_PROVIDER_SOURCE" \
        --terraform-provider-version="$TERRAFORM_PROVIDER_VERSION"
else
    echo "Starting provider-palette in NORMAL mode with verbose logging..."
    echo "Provider version: $TERRAFORM_PROVIDER_VERSION"
    echo "Logs will be shown in real-time"
    echo "Use './debug.sh debug' to run in Terraform debugging mode"
    echo ""
    
    # Run the Crossplane provider with verbose logging (--debug flag)
    go run cmd/provider/main.go \
        --debug \
        --enable-management-policies \
        --terraform-version="$TERRAFORM_VERSION" \
        --terraform-provider-source="$TERRAFORM_PROVIDER_SOURCE" \
        --terraform-provider-version="$TERRAFORM_PROVIDER_VERSION"
fi 