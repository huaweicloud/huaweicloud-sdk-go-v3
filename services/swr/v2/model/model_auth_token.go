package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthToken struct {

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	// 过期时间
	ExpireDate *string `json:"expire_date,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 访问凭证名称
	Name *string `json:"name,omitempty"`

	// 访问凭证ID
	TokenId *string `json:"token_id,omitempty"`

	// 访问凭证用户名
	UserId *string `json:"user_id,omitempty"`

	// 用户访问凭据信息
	UserProfile *string `json:"user_profile,omitempty"`
}

func (o AuthToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthToken struct{}"
	}

	return strings.Join([]string{"AuthToken", string(data)}, " ")
}
