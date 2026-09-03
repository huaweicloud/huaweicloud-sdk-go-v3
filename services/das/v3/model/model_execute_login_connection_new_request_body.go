package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteLoginConnectionNewRequestBody 登录操作请求体
type ExecuteLoginConnectionNewRequestBody struct {
	Login *LoginInfo `json:"login,omitempty"`

	Logout *LogoutInfo `json:"logout,omitempty"`

	RetryLogin *RetryLoginInfo `json:"retry_login,omitempty"`
}

func (o ExecuteLoginConnectionNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteLoginConnectionNewRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteLoginConnectionNewRequestBody", string(data)}, " ")
}
