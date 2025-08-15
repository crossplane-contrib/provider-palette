package cluster

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures cluster-scoped resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Add cluster-specific configurations here if needed
	// Most configurations are handled in common/config.go
}
