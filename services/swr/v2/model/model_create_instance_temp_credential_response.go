package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceTempCredentialResponse Response Object
type CreateInstanceTempCredentialResponse struct {

	// 临时访问凭证密码
	AuthToken *string `json:"auth_token,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 过期时间
	ExpireDate *string `json:"expire_date,omitempty"`

	// 临时访问凭证ID
	TokenId *string `json:"token_id,omitempty"`

	// 临时访问凭证的用户ID
	UserId         *string `json:"user_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceTempCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceTempCredentialResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceTempCredentialResponse", string(data)}, " ")
}
