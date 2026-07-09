// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositorycollaborators


type RepositoryCollaboratorsTeam struct {
	// ID or slug of the team to add to the repository as a collaborator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_collaborators#team_id RepositoryCollaborators#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// Permission to grant to the team.
	//
	// Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization. Defaults to `push`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_collaborators#permission RepositoryCollaborators#permission}
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

