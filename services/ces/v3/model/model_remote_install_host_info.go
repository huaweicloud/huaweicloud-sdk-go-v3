package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoteInstallHostInfo struct {

	// 被远程安装的主机名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 被远程安装的主机IP
	RemoteIp *string `json:"remote_ip,omitempty"`

	// 被远程安装的主机的登录用户名
	UserName *string `json:"user_name,omitempty"`

	// 被远程安装的主机的登录端口
	Port *string `json:"port,omitempty"`

	// 被远程安装的主机的登录密码
	Password *string `json:"password,omitempty"`

	// 被远程安装的主机远程连接是否采用秘钥方式（false时为密码方式）
	RemoteUsePem *bool `json:"remote_use_pem,omitempty"`
}

func (o RemoteInstallHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteInstallHostInfo struct{}"
	}

	return strings.Join([]string{"RemoteInstallHostInfo", string(data)}, " ")
}
