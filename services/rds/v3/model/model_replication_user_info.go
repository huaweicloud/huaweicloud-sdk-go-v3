package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplicationUserInfo 发布订阅用户登录信息。
type ReplicationUserInfo struct {

	// 服务器ip。
	ServerIp string `json:"server_ip"`

	// 端口号。
	ServerPort int32 `json:"server_port"`

	// 服务器名称。
	ServerName string `json:"server_name"`

	// 登录名。
	LoginUserName string `json:"login_user_name"`

	// 登录密码。要求密码长度在5~64位之间。
	LoginUserPassword string `json:"login_user_password"`
}

func (o ReplicationUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationUserInfo struct{}"
	}

	return strings.Join([]string{"ReplicationUserInfo", string(data)}, " ")
}
