package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateUserReq struct {

	// 用户名称。
	UserName string `json:"user_name"`

	// 用户邮箱。
	UserEmail *string `json:"user_email,omitempty"`

	// 账户过期时间，0表示永远不过期。时间格式：yyyy-MM-ddTHH:mm:ssZ。
	AccountExpires *string `json:"account_expires,omitempty"`

	// 是否允许用户更改密码，缺省值为true。
	EnableChangePassword *bool `json:"enable_change_password,omitempty"`

	// 下次登录是否必须更改密码，缺省值为true。
	NextLoginChangePassword *bool `json:"next_login_change_password,omitempty"`

	// 用户组的专有ID列表。
	GroupIds *[]string `json:"group_ids,omitempty"`

	// 用户描述，字符串长度区间[0, 255]。
	Description *string `json:"description,omitempty"`

	// 别名。
	AliasName *string `json:"alias_name,omitempty"`
}

func (o CreateUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserReq struct{}"
	}

	return strings.Join([]string{"CreateUserReq", string(data)}, " ")
}
