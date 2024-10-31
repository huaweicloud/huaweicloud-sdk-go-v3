package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// User 租户信息
type User struct {

	// 账户名称
	DomainName *string `json:"domain_name,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}
