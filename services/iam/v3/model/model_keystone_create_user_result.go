package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeystoneCreateUserResult struct {
	// IAM用户所属账号ID。

	DomainId string `json:"domain_id"`
	// IAM用户名。

	Name string `json:"name"`
	// IAM用户描述信息。

	Description *string `json:"description,omitempty"`
	// IAM用户密码过期时间（UTC时间），“null”表示密码不过期。

	PasswordExpiresAt string `json:"password_expires_at"`

	Links *LinksSelf `json:"links"`
	// IAM用户ID。

	Id string `json:"id"`
	// IAM用户是否启用。true表示启用，false表示停用，默认为true。

	Enabled bool `json:"enabled"`
}

func (o KeystoneCreateUserResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserResult struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserResult", string(data)}, " ")
}
