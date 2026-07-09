// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type GithubProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#alias GithubProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// app_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#app_auth GithubProvider#app_auth}
	AppAuth *GithubProviderAppAuth `field:"optional" json:"appAuth" yaml:"appAuth"`
	// The base URL for the GitHub API;
	//
	// this defaults to the GitHub API URL. If you are using GitHub Enterprise Server (GHES) or GitHub Enterprise Cloud with Data Residency (GHEC-DR), this is required. This can also be set by the `GITHUB_BASE_URL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#base_url GithubProvider#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// The path to the cache directory for persisting GitHub API requests between runs;
	//
	// if not set there will be no caching between runs. This can also be set by the `GITHUB_CACHE_PATH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#cache_path GithubProvider#cache_path}
	CachePath *string `field:"optional" json:"cachePath" yaml:"cachePath"`
	// Allow insecure server connections when using SSL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#insecure GithubProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// Use the legacy GitHub client implementation;
	//
	// if set to `false`, the new client implementation is used. This can also be set by the `GITHUB_LEGACY_CLIENT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#legacy_client GithubProvider#legacy_client}
	LegacyClient interface{} `field:"optional" json:"legacyClient" yaml:"legacyClient"`
	// The maximum number of results per page for paginated API requests;
	//
	// this defaults to `100`. This can also be set by the `GITHUB_MAX_PER_PAGE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#max_per_page GithubProvider#max_per_page}
	MaxPerPage *float64 `field:"optional" json:"maxPerPage" yaml:"maxPerPage"`
	// The maximum number of retries for failed requests; this defaults to `3`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#max_retries GithubProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// GitHub organization to manage. This can also be set by the `GITHUB_ORGANIZATION` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#organization GithubProvider#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// GitHub organization or user account to manage;
	//
	// this is required when authenticating using a GitHub App. If the owner is not provided and a token is provided, the provider will attempt to auto-detect the owner associated with the token. This can also be set by the `GITHUB_OWNER` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#owner GithubProvider#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Allow the provider to make parallel API calls;
	//
	// this is experimental and may cause concurrency and rate limiting issues. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation is designed to safely handle parallel requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#parallel_requests GithubProvider#parallel_requests}
	ParallelRequests interface{} `field:"optional" json:"parallelRequests" yaml:"parallelRequests"`
	// The delay in milliseconds between read operations;
	//
	// this defaults to `0`. This can be used to mitigate rate limiting issues when performing a large number of read operations. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation is GitHub rate limit aware.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#read_delay_ms GithubProvider#read_delay_ms}
	ReadDelayMs *float64 `field:"optional" json:"readDelayMs" yaml:"readDelayMs"`
	// List of HTTP status codes that should be retried;
	//
	// if not set this uses the provider defaults. This setting only applies when `max_retries` is greater than `0`. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation handles the retry logic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#retryable_errors GithubProvider#retryable_errors}
	RetryableErrors *[]*float64 `field:"optional" json:"retryableErrors" yaml:"retryableErrors"`
	// The delay in milliseconds between retry attempts;
	//
	// this defaults to `1000`. This setting only applies when `max_retries` is greater than `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#retry_delay_ms GithubProvider#retry_delay_ms}
	RetryDelayMs *float64 `field:"optional" json:"retryDelayMs" yaml:"retryDelayMs"`
	// GitHub OAuth or Personal Access Token (PAT) to use for authentication.
	//
	// This can also be set by the `GITHUB_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#token GithubProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The delay in milliseconds between write operations;
	//
	// this defaults to `1000`. This is used to mitigate the GitHub API's abuse rate limits when writing. Note that **ALL** requests to the GraphQL API are implemented as `POST` requests under the hood, so this setting affects those calls as well. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation is GitHub rate limit aware.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#write_delay_ms GithubProvider#write_delay_ms}
	WriteDelayMs *float64 `field:"optional" json:"writeDelayMs" yaml:"writeDelayMs"`
}

