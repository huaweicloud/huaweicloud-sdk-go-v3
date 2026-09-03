package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginInfo 已保存密码的登录信息
type LoginInfo struct {

	// 登录数据库的用户名
	Username *string `json:"username,omitempty"`
}

func (o LoginInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginInfo struct{}"
	}

	return strings.Join([]string{"LoginInfo", string(data)}, " ")
}
