package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type LoginTokenUser struct {
	Domain *LoginTokenDomain `json:"domain,omitempty" xml:"domain"`

	// 被委托方用户名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 被委托方用户的密码过期时间。
	PasswordExpiresAt *string `json:"password_expires_at,omitempty" xml:"password_expires_at"`

	// 被委托方用户ID。
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o LoginTokenUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginTokenUser struct{}"
	}

	return strings.Join([]string{"LoginTokenUser", string(data)}, " ")
}
