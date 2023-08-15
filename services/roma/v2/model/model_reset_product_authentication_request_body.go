package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetProductAuthenticationRequestBody struct {

	// 产品用户名，支持英文大小写、英文符号(-)及数字，长度限制10-50位，传参空，用户名不被重置
	UserName *string `json:"user_name,omitempty"`

	// 产品密码，包含数字、英文字母大小写、特殊字符(-~!@#$%^&*()-_=+|[{}];:<>/?)，长度10-80位，传参空，密码不被重置。当用户名与密码都为空时，密码重置，由系统生成。
	Password *string `json:"password,omitempty"`
}

func (o ResetProductAuthenticationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetProductAuthenticationRequestBody struct{}"
	}

	return strings.Join([]string{"ResetProductAuthenticationRequestBody", string(data)}, " ")
}
