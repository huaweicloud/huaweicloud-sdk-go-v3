package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LocalImageHostInfo struct {

	// Agent ID
	AgentId *string `json:"agent_id,omitempty"`

	// agent_status
	AgentStatus *string `json:"agent_status,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器开通的版本
	Version *string `json:"version,omitempty"`

	// 私有IP地址
	PrivateIp *string `json:"private_ip,omitempty"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// docker名称
	DockerName *string `json:"docker_name,omitempty"`

	// docker类型
	DockerType *string `json:"docker_type,omitempty"`
}

func (o LocalImageHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalImageHostInfo struct{}"
	}

	return strings.Join([]string{"LocalImageHostInfo", string(data)}, " ")
}
