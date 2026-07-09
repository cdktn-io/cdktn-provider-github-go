// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositorycollaborators


type RepositoryCollaboratorsUser struct {
	// Login for the user to add to the repository as a collaborator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_collaborators#username RepositoryCollaborators#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Permission to grant to the user.
	//
	// Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization. Must be `push` for personal repositories. Defaults to `push`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_collaborators#permission RepositoryCollaborators#permission}
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

