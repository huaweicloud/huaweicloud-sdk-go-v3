package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifySupabaseApplicationAuthConfigRequestBody Auth配置修改请求。部分更新：仅更新非空字段，省略=不更新，传null=保留。 @JsonInclude NON_EMPTY。
type ModifySupabaseApplicationAuthConfigRequestBody struct {
	EnableAccountSignUp *bool `json:"enable_account_sign_up,omitempty"`

	EnableAnonymousUsers *bool `json:"enable_anonymous_users,omitempty"`

	ApiExternalUrl *string `json:"api_external_url,omitempty"`

	SiteUrl *string `json:"site_url,omitempty"`

	// 逗号分隔URL，最多10个。
	AllowRedirectUrls *string `json:"allow_redirect_urls,omitempty"`

	GithubOauthProvider *SupabaseGitHubOAuthProviderReq `json:"github_oauth_provider,omitempty"`

	EmailOauthProvider *SupabaseEmailOAuthProviderReq `json:"email_oauth_provider,omitempty"`

	EmailTemplateConfig *SupabaseEmailTemplateReq `json:"email_template_config,omitempty"`
}

func (o ModifySupabaseApplicationAuthConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySupabaseApplicationAuthConfigRequestBody struct{}"
	}

	return strings.Join([]string{"ModifySupabaseApplicationAuthConfigRequestBody", string(data)}, " ")
}
