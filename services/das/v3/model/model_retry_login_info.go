package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryLoginInfo 未保存密码的登录信息
type RetryLoginInfo struct {

	// 登录数据库的用户名
	Username *string `json:"username,omitempty"`

	// 登录数据库的密码
	Password string `json:"password"`

	// 是否保存密码
	IsSavePassword bool `json:"is_save_password"`
}

func (o RetryLoginInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryLoginInfo struct{}"
	}

	return strings.Join([]string{"RetryLoginInfo", string(data)}, " ")
}
