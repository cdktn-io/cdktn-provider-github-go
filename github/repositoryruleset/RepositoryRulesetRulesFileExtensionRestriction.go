// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesFileExtensionRestriction struct {
	// A list of file extensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.0/docs/resources/repository_ruleset#restricted_file_extensions RepositoryRuleset#restricted_file_extensions}
	RestrictedFileExtensions *[]*string `field:"required" json:"restrictedFileExtensions" yaml:"restrictedFileExtensions"`
}

