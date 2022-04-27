package aws_infra_config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InfrastructureConfig struct {
	metav1.TypeMeta `json:",inline"`

	// EnableECRAccess specifies whether the IAM role policy for the worker nodes shall contain
	// permissions to access the ECR.
	// default: true
	// +optional
	EnableECRAccess *bool `json:"enableECRAccess,omitempty"`

	// Networks is the AWS specific network configuration (VPC, subnets, etc.)
	Networks Networks `json:"networks"`

	// IgnoreTags allows to configure which resource tags on resources managed by Gardener should be ignored during
	// infrastructure reconciliation. By default, all tags that are added outside of Gardener's / terraform's
	// reconciliation will be removed during the next reconciliation. This field allows users and automation to add
	// custom tags on resources created and managed by Gardener without loosing them on the next reconciliation.
	// See https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/resource-tagging#ignoring-changes-in-all-resources
	// for details of the underlying terraform implementation.
	// +optional
	IgnoreTags *IgnoreTags `json:"ignoreTags,omitempty"`
}
