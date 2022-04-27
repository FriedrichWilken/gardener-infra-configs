package azure_infra_config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InfrastructureConfig struct {
	metav1.TypeMeta `json:",inline"`
	// ResourceGroup is azure resource group.
	// +optional
	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`
	// Networks is the network configuration (VNet, subnets, etc.).
	Networks NetworkConfig `json:"networks"`
	// Identity contains configuration for the assigned managed identity.
	// +optional
	Identity *IdentityConfig `json:"identity,omitempty"`
	// Zoned indicates whether the cluster uses availability zones.
	// +optional
	Zoned bool `json:"zoned,omitempty"`
}
