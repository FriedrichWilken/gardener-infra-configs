package gcp_infra_config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InfrastructureConfig infrastructure configuration resource
type InfrastructureConfig struct {
	metav1.TypeMeta `json:",inline"`

	// Networks is the network configuration (VPC, subnets, etc.)
	Networks NetworkConfig `json:"networks"`
}
