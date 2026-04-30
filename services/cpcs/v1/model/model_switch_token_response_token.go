package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchTokenResponseToken 换取的token实例
type SwitchTokenResponseToken struct {
	App *SwitchTokenResponseTokenApp `json:"app,omitempty"`

	// 过期时间
	ExpiresAt *string `json:"expires_at,omitempty"`

	// 角色列表
	Roles *[]SwitchTokenResponseTokenRoles `json:"roles,omitempty"`

	// 签发时间
	IssuedAt *string `json:"issued_at,omitempty"`

	User *SwitchTokenResponseTokenUser `json:"user,omitempty"`
}

func (o SwitchTokenResponseToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchTokenResponseToken struct{}"
	}

	return strings.Join([]string{"SwitchTokenResponseToken", string(data)}, " ")
}
