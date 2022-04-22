package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProductAuthenticationResponse struct {

	// 一型一密/一机一密的用户名
	UserName *string `json:"user_name,omitempty"`

	// 一型一密/一机一密的密码，输入要求：至少1个数字，1个大写字母，1个小写字母，1个特殊字符(~!@#$%^&*()-_=+|[{}];:<>/?)，长度8-32个字符
	Password       *string `json:"password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProductAuthenticationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProductAuthenticationResponse struct{}"
	}

	return strings.Join([]string{"ShowProductAuthenticationResponse", string(data)}, " ")
}
