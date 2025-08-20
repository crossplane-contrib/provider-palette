# Palette Provider Crossplane Composition Example

This composition demonstrates how to use Crossplane with the Palette provider to create a complete customer onboarding workflow, including project creation, cloud account provisioning, and EKS cluster deployment.

## File Structure

```
composition/
├── README.md                        # This documentation file
├── provider.yaml                    # Palette provider installation
├── spectrocloud-xp-ns.yaml          # Namespace for Crossplane resources
├── provider-config.yaml             # Provider configuration for Palette API
├── palette-creds.yaml               # Secret containing Palette API credentials
└── composite-resources/
    ├── composition.yaml             # Main composition with pipeline steps
    ├── function.yaml                # Required Crossplane functions
    ├── rbac.yaml                    # RBAC permissions for namespace management
    ├── xr.yaml                      # Example XCustomer resource
    └── xrd.yaml                     # Composite Resource Definition for XCustomer
```

## Overview

The `XCustomer` composition creates a complete customer environment in Palette by orchestrating multiple resources across three main pipeline steps:

1. **Onboarding** - Creates customer project, namespace, and provider configuration
2. **Cloud Account Provisioning** - Sets up AWS cloud account with proper credentials
3. **Cluster Provisioning** - Deploys an EKS cluster with specified configuration

## Pipeline Steps

### Step 1: Onboarding

Creates the foundational resources needed for customer isolation and provider access.

#### Resources Created:
- **CustomerProject** - New project in Palette (created in `spectrocloud-xp` namespace)
- **CustomerNamespace** - Dedicated namespace for customer resources
- **CustomerProviderSecret** - API credentials for the customer's project
- **CustomerProviderConfig** - Provider configuration using customer credentials
- **SecretUsesNamespace** - Protection mechanism to prevent premature deletion
- **ProviderConfigUsesSecret** - Protection mechanism for provider config

#### Dependencies:
```
CustomerProject (spectrocloud-xp namespace)
    ↓
CustomerNamespace
    ↓
CustomerProviderSecret
    ↓
CustomerProviderConfig
```

### Step 2: Cloud Account Provisioning

Sets up AWS cloud account access for the customer.

#### Resources Created:
- **AWSCloudAccountSecret** - Secret containing AWS credentials
- **CloudAccount** - AWS cloud account configuration
- **CloudAccountUsesAWSCloudAccountSecret** - Protection mechanism

#### Dependencies:
```
AWSCloudAccountSecret
    ↓
CloudAccount
```

### Step 3: Cluster Provisioning

Deploys an EKS cluster with the specified configuration.

#### Resources Created:
- **ProvisionClusterEKS** - EKS cluster with customer configuration
- **ClusterUsesProviderConfig** - Protection mechanism

#### Dependencies:
```
ProvisionClusterEKS
    ↓ (depends on CloudAccount and CustomerProviderConfig)
```

## Resource Protection and Deletion Order

The composition uses Crossplane's protection mechanisms to ensure proper resource deletion order:

```
ProvisionClusterEKS
    ↓ (Usage: ClusterUsesProviderConfig)
CustomerProviderConfig
    ↓ (Usage: ProviderConfigUsesSecret)
CustomerProviderSecret
    ↓ (Usage: SecretUsesNamespace)
CustomerNamespace
    ↓
CloudAccount
    ↓ (Usage: CloudAccountUsesAWSCloudAccountSecret)
AWSCloudAccountSecret
    ↓
CustomerProject
```

## XCustomer Resource Specification

The `XCustomer` composite resource accepts the following configuration:

### Required Fields:
- `name` - Customer identifier
- `host` - Palette console host
- `apiKey` - Palette API key

### Optional Fields:
- `awsAccessKey` - AWS access key for cloud account
- `awsSecretKey` - AWS secret key for cloud account
- `clusterName` - Name for the EKS cluster
- `clusterDescription` - Description for the cluster
- `clusterTags` - Tags to apply to the cluster
- `clusterProfileID` - Palette cluster profile ID
- `clusterConfig` - Cluster configuration details
  - `cloudConfig` - Cloud-specific settings (region, SSH key, access control)
  - `workerPoolsConfig` - Worker node pool configurations

## Usage Example

1. **Install the Palette provider**:
   ```bash
   kubectl apply -f provider.yaml
   ```

2. **Create the spectrocloud-xp namespace**:
   ```bash
   kubectl apply -f spectrocloud-xp-ns.yaml
   ```

3. **Apply the provider configuration and credentials**:
   ```bash
   kubectl apply -f palette-creds.yaml
   kubectl apply -f provider-config.yaml
   ```

4. **Install the required functions and RBAC**:
   ```bash
   kubectl apply -f composite-resources/function.yaml
   kubectl apply -f composite-resources/rbac.yaml
   ```

5. **Create the composite resource definition**:
   ```bash
   kubectl apply -f composite-resources/xrd.yaml
   ```

6. **Apply the composition**:
   ```bash
   kubectl apply -f composite-resources/composition.yaml
   ```

7. **Create a customer instance**:
   ```bash
   kubectl apply -f composite-resources/xr.yaml
   ```

## Notes

- The composition uses the `crossplane-contrib-function-patch-and-transform` function for resource creation
- All customer resources are created in a dedicated namespace for isolation
- The project is created in the `spectrocloud-xp` namespace using a tenant-level provider configuration
- AWS credentials are stored as Kubernetes secrets and referenced by the cloud account
- The EKS cluster configuration supports single worker pool configuration (limitation noted in TODO comment)
- Protection mechanisms ensure resources are deleted in the correct order to prevent dependency conflicts
