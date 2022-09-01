package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ScopedTokenUser struct {
	Domain *TokenDomainResult `json:"domain" xml:"domain"`

	OsFederation *TokenUserOsfederation `json:"OS-FEDERATION" xml:"OS-FEDERATION"`

	// 用户ID。
	Id string `json:"id" xml:"id"`

	// 用户名。
	Name string `json:"name" xml:"name"`

	// 密码过期时间（UTC时间），“”表示密码不过期。
	PasswordExpiresAt string `json:"password_expires_at" xml:"password_expires_at"`
}

func (o ScopedTokenUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScopedTokenUser struct{}"
	}

	return strings.Join([]string{"ScopedTokenUser", string(data)}, " ")
}
