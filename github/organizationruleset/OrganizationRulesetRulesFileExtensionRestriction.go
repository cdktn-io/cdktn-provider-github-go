// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesFileExtensionRestriction struct {
	// The file extensions that are restricted from being pushed to the commit graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#restricted_file_extensions OrganizationRuleset#restricted_file_extensions}
	RestrictedFileExtensions *[]*string `field:"required" json:"restrictedFileExtensions" yaml:"restrictedFileExtensions"`
}

