package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseGitHubOAuthProviderReq enable_github=false时清空oauth_client_id/secret/redirect_uri。
type SupabaseGitHubOAuthProviderReq struct {
	EnableGithub *bool `json:"enable_github,omitempty"`

	OauthClientId *string `json:"oauth_client_id,omitempty"`

	// 秘密，受控输入，show不返回。
	OauthClientSecret *string `json:"oauth_client_secret,omitempty"`

	RedirectUri *string `json:"redirect_uri,omitempty"`
}

func (o SupabaseGitHubOAuthProviderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseGitHubOAuthProviderReq struct{}"
	}

	return strings.Join([]string{"SupabaseGitHubOAuthProviderReq", string(data)}, " ")
}
