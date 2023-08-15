package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetAuthenticationRequestBody 设备用户密码更新请求
type ResetAuthenticationRequestBody struct {

	// 设备用户名，支持英文大小写、英文符号(-)及数字，长度限制10-50位，传参空，用户名不被重置
	UserName *string `json:"user_name,omitempty"`

	// 设备密码，输入要求：至少1数字，1大写字母，1小写字母，1特殊字符(~!@#$%^&*()-_=+|[{}];:<>/?)，长度8-32位，传参空，密码不被重置。当用户名与密码都为空时，密码重置，由系统生成。
	Password *string `json:"password,omitempty"`
}

func (o ResetAuthenticationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetAuthenticationRequestBody struct{}"
	}

	return strings.Join([]string{"ResetAuthenticationRequestBody", string(data)}, " ")
}
