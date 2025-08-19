# Palette Provider Crossplane Composition Example

## File Structure

```
composition/
├── README.md                        # This documentation file
├── provider.yaml                    # Palette provider installation
├── spectrocloud-xp-ns.yaml          # Namespace for Crossplane resources
├── provider-config.yaml             # Provider configuration for Palette API
├── palette-creds.yaml               # Secret containing Palette API credentials
└── composite-resources/
    ├── composition.yaml               # Main composition with pipeline steps and usage resources
    ├── function.yaml                  # Required Crossplane functions
    ├── rbac.yaml                      # RBAC permissions for namespace management
    ├── xr.yaml                        # Example XCustomer resource
    └── xrd.yaml                       # Composite Resource Definition for XCustomer
```

## Resource Creation Flow

This composition creates resources in a specific dependency order to ensure proper setup and includes protection mechanisms to prevent premature deletion.

### Dependency Graph

#### Resource Creation Order
```
CustomerProject
    ↓
CustomerNamespace
    ↓
CustomerProviderSecret
    ↓
CustomerProviderConfig
    ↓
ProvisionClusterEKS
```

### Resource Creation Steps

1. **CustomerProject** - Provisions a new project in Palette using the customer's specifications
2. **CustomerNamespace** - Creates a dedicated namespace for the customer's resources
3. **CustomerProviderSecret** - Generates API credentials specific to the customer's project
4. **CustomerProviderConfig** - Configures the Palette provider with the customer-specific credentials
5. **ProvisionClusterEKS** - Deploys an EKS cluster within the customer's project using the configured provider

#### Resource Deletion Order (Protected by Usage Resources)
```
ProvisionClusterEKS
    ↓ (Usage: ClusterUsesProviderConfig)
CustomerProviderConfig
    ↓ (Usage: ProviderConfigUsesSecret)
CustomerProviderSecret
    ↓ (Usage: SecretUsesNamespace)
CustomerNamespace
    ↓ 
CustomerProject
```

Each step depends on the successful completion of the previous step, ensuring resources are created in the correct order and with proper dependencies.

### Usage Resource Protection Details

The composition includes inline Usage resources within each pipeline step to enforce proper deletion order using Crossplane's protection mechanisms:

- **`ClusterUsesProviderConfig`** - Prevents ProviderConfig deletion while EKS cluster exists
- **`ProviderConfigUsesSecret`** - Prevents Secret deletion while ProviderConfig exists  
- **`SecretUsesNamespace`** - Prevents Namespace deletion while Secret exists

These Usage resources are defined inline within the composition pipeline steps, ensuring that when the composition is deleted, resources are removed in the reverse order of creation, preventing dependency conflicts and orphaned resources.
