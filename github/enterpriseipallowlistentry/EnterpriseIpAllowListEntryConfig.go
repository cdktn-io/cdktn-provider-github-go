// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package enterpriseipallowlistentry

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EnterpriseIpAllowListEntryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The slug of the enterprise to apply the IP allow list entry to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/enterprise_ip_allow_list_entry#enterprise_slug EnterpriseIpAllowListEntry#enterprise_slug}
	EnterpriseSlug *string `field:"required" json:"enterpriseSlug" yaml:"enterpriseSlug"`
	// An IP address or range of IP addresses in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/enterprise_ip_allow_list_entry#ip EnterpriseIpAllowListEntry#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/enterprise_ip_allow_list_entry#id EnterpriseIpAllowListEntry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the entry is currently active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/enterprise_ip_allow_list_entry#is_active EnterpriseIpAllowListEntry#is_active}
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// An optional name for the IP allow list entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/enterprise_ip_allow_list_entry#name EnterpriseIpAllowListEntry#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

