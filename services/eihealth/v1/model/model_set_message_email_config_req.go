package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 消息邮件发送配置请求体
type SetMessageEmailConfigReq struct {

	// 服务器地址
	Server string `json:"server"`

	// 展示名
	SubjectPrefix *string `json:"subject_prefix,omitempty"`

	// 用户名
	UserName string `json:"user_name"`

	// 密码
	Password string `json:"password"`

	// 邮箱
	Email string `json:"email"`

	Language *LanguageEnum `json:"language"`
}

func (o SetMessageEmailConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMessageEmailConfigReq struct{}"
	}

	return strings.Join([]string{"SetMessageEmailConfigReq", string(data)}, " ")
}
