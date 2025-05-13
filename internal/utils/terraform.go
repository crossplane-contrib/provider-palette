package utils

import (
	"fmt"
	"os"
)

// ConfigureTFReattachProvider configures the TF_REATTACH_PROVIDERS environment variable
func ConfigureTFReattachProvider() error {
	addr := os.Getenv("TF_ADDR")
	pid := os.Getenv("TF_PID")
	if addr == "" || pid == "" {
		return nil
	}
	reattachProvider := fmt.Sprintf(`{
		"registry.terraform.io/spectrocloud/spectrocloud": {
			"Protocol": "grpc",
			"ProtocolVersion": 5,
			"Pid": %s,
			"Test": true,
			"Addr": {
				"Network": "unix",
				"String": "%s"
			}
		}
	}`, pid, addr)
	return os.Setenv("TF_REATTACH_PROVIDERS", reattachProvider)
}
