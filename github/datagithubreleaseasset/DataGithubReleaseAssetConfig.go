// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datagithubreleaseasset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGithubReleaseAssetConfig struct {
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
	// ID of the release asset to retrieve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/data-sources/release_asset#asset_id DataGithubReleaseAsset#asset_id}
	AssetId *float64 `field:"required" json:"assetId" yaml:"assetId"`
	// Owner of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/data-sources/release_asset#owner DataGithubReleaseAsset#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Name of the repository to retrieve the release asset from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/data-sources/release_asset#repository DataGithubReleaseAsset#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Whether to download the asset file content into the `file_contents` attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/data-sources/release_asset#download_file_contents DataGithubReleaseAsset#download_file_contents}
	DownloadFileContents interface{} `field:"optional" json:"downloadFileContents" yaml:"downloadFileContents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/data-sources/release_asset#id DataGithubReleaseAsset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

