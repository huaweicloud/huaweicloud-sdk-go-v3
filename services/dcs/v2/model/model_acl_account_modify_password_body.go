package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclAccountModifyPasswordBody struct {

	// 旧密码
	OldPassword *string `json:"old_password,omitempty"`

	// 新密码 - 输入长度为8到64位的字符串。 - 不能包含正序逆序用户名。 - 必须包含如下四种字符中的三种组合（不允许包含冒号）：   - 小写字母   - 大写字母   - 数字   - 特殊字符包括（`~!@#$%^&*()-_=+\\|{},<.>/?）
	NewPassword *string `json:"new_password,omitempty"`
}

func (o AclAccountModifyPasswordBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclAccountModifyPasswordBody struct{}"
	}

	return strings.Join([]string{"AclAccountModifyPasswordBody", string(data)}, " ")
}
