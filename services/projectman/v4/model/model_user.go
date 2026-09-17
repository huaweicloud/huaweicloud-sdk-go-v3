package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// User 项目管理用户
type User struct {

	// 性别
	Gender *string `json:"gender,omitempty"`

	// 用户uuid
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`

	// 租户uuid
	DomainId *string `json:"domain_id,omitempty"`

	// 租户名
	DomainName *string `json:"domain_name,omitempty"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}
