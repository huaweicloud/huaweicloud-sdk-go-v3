package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 邮箱服务器连通性测试
type TryEmailConnectionReq struct {

	// 服务器地址
	Server string `json:"server"`

	// 用户名
	UserName string `json:"user_name"`

	// 密码
	Password string `json:"password"`

	// 邮箱
	Email string `json:"email"`

	Language *LanguageEnum `json:"language"`
}

func (o TryEmailConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TryEmailConnectionReq struct{}"
	}

	return strings.Join([]string{"TryEmailConnectionReq", string(data)}, " ")
}
