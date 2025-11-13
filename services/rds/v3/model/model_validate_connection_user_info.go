package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateConnectionUserInfo 连接测试用户登录信息。
type ValidateConnectionUserInfo struct {

	// 服务器ip。
	ServerIp *string `json:"server_ip,omitempty"`

	// 端口号。
	ServerPort *int32 `json:"server_port,omitempty"`

	// 登录名。
	LoginUserName *string `json:"login_user_name,omitempty"`

	// 登录密码。要求密码长度在5~64位之间。
	LoginUserPassword *string `json:"login_user_password,omitempty"`
}

func (o ValidateConnectionUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConnectionUserInfo struct{}"
	}

	return strings.Join([]string{"ValidateConnectionUserInfo", string(data)}, " ")
}
