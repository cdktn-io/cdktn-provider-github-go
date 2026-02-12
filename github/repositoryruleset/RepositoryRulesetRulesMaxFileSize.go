// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesMaxFileSize struct {
	// The maximum allowed size of a file in megabytes (MB). Valid range is 1-100 MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/repository_ruleset#max_file_size RepositoryRuleset#max_file_size}
	MaxFileSize *float64 `field:"required" json:"maxFileSize" yaml:"maxFileSize"`
}

