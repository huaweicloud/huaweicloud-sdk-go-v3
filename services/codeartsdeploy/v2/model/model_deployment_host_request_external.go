package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentHostRequestExternal struct {

	// 是否为代理主机
	AsProxy *bool `json:"as_proxy,omitempty"`

	Authorization *DeploymentHostAuthorizationBody `json:"authorization,omitempty"`

	// 主机名称
	HostName string `json:"host_name"`

	// 主机ip，如：161.17.101.12
	Ip string `json:"ip"`

	// ssh端口，如：22
	Port int32 `json:"port"`

	// 代理主机id
	ProxyHostId *string `json:"proxy_host_id,omitempty"`

	// 是否同步主机信息
	Sync *bool `json:"sync,omitempty"`

	// 是否安装icAgent
	InstallIcagent *bool `json:"install_icagent,omitempty"`
}

func (o DeploymentHostRequestExternal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentHostRequestExternal struct{}"
	}

	return strings.Join([]string{"DeploymentHostRequestExternal", string(data)}, " ")
}
