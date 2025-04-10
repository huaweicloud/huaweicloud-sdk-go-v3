package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginWhiteListRequestInfo 登录白名单
type LoginWhiteListRequestInfo struct {

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 登录源IP
	LoginIp *string `json:"login_ip,omitempty"`

	// 登录用户名
	LoginUserName *string `json:"login_user_name,omitempty"`
}

func (o LoginWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"LoginWhiteListRequestInfo", string(data)}, " ")
}
