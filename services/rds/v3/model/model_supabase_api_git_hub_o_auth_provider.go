package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseApiGitHubOAuthProvider 查询响应不含oauth_client_secret。
type SupabaseApiGitHubOAuthProvider struct {
	EnableGithub *bool `json:"enable_github,omitempty"`

	OauthClientId *string `json:"oauth_client_id,omitempty"`

	RedirectUri *string `json:"redirect_uri,omitempty"`
}

func (o SupabaseApiGitHubOAuthProvider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseApiGitHubOAuthProvider struct{}"
	}

	return strings.Join([]string{"SupabaseApiGitHubOAuthProvider", string(data)}, " ")
}
