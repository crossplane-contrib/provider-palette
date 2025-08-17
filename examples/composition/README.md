# Palette Provider Crossplane Composition Example

## File Structure

```
composition/
├── README.md                          # This documentation file
├── provider.yaml                      # Palette provider installation
├── spectrocloud-xp-ns.yaml           # Namespace for Crossplane resources
├── provider-config.yaml              # Provider configuration for Palette API
├── palette-creds.yaml                # Secret containing Palette API credentials
└── composite-resources/
    ├── composition.yaml               # Main composition with pipeline steps
    ├── function.yaml                  # Required Crossplane functions
    ├── rbac.yaml                      # RBAC permissions for namespace management
    ├── xr.yaml                        # Example XCustomer resource
    └── xrd.yaml                       # Composite Resource Definition for XCustomer
```