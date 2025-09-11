package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoteInstallHostInfo struct {

	// **参数解释**: 被远程安装的主机名称 **取值范围**: 数组长度范围为[1,128]
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**: 被远程安装的主机IP **取值范围**: 仅由数字(0-9)和小数点(.)组成的字符串，字符串长度为[1,15]
	RemoteIp *string `json:"remote_ip,omitempty"`

	// **参数解释**: 被远程安装的主机的登录用户名 **取值范围**: 数组长度范围为[1,16]
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 被远程安装的主机的登录端口 **取值范围**: 数组长度范围为[1,5]
	Port *string `json:"port,omitempty"`

	// **参数解释**: 被远程安装的主机的登录密码 **取值范围**: 数组长度范围为[1,3000]
	Password *string `json:"password,omitempty"`

	// **参数解释**: 被远程安装的主机远程连接是否采用密钥方式（false时为密码方式） **取值范围**: - false: 密码方式 - true: 密钥方式
	RemoteUsePem *bool `json:"remote_use_pem,omitempty"`
}

func (o RemoteInstallHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteInstallHostInfo struct{}"
	}

	return strings.Join([]string{"RemoteInstallHostInfo", string(data)}, " ")
}
