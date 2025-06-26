package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceLtCredentialResponse Response Object
type CreateInstanceLtCredentialResponse struct {

	// 访问凭证密码
	AuthToken *string `json:"auth_token,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 过期时间
	ExpireDate *string `json:"expire_date,omitempty"`

	// 访问凭证ID
	TokenId *string `json:"token_id,omitempty"`

	// 访问凭证用户名
	UserId         *string `json:"user_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceLtCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceLtCredentialResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceLtCredentialResponse", string(data)}, " ")
}
