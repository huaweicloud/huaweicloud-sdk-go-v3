package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSupabaseApplicationAuthConfigResponse Response Object
type ShowSupabaseApplicationAuthConfigResponse struct {
	UpdatedTime *int64 `json:"updated_time,omitempty"`

	EnableAccountSignUp *bool `json:"enable_account_sign_up,omitempty"`

	EnableAnonymousUsers *bool `json:"enable_anonymous_users,omitempty"`

	ApiExternalUrl *string `json:"api_external_url,omitempty"`

	SiteUrl *string `json:"site_url,omitempty"`

	AllowRedirectUrls *string `json:"allow_redirect_urls,omitempty"`

	GithubOauthProvider *SupabaseApiGitHubOAuthProvider `json:"github_oauth_provider,omitempty"`

	EmailOauthProvider *SupabaseApiEmailOAuthProvider `json:"email_oauth_provider,omitempty"`

	EmailTemplateConfig *SupabaseApiEmailTemplate `json:"email_template_config,omitempty"`
	HttpStatusCode      int                       `json:"-"`
}

func (o ShowSupabaseApplicationAuthConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupabaseApplicationAuthConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowSupabaseApplicationAuthConfigResponse", string(data)}, " ")
}
