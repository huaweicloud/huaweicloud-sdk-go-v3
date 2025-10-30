package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LocalImageContainerInfo struct {

	// Agent id
	AgentId *string `json:"agent_id,omitempty"`

	// 容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// 容器ID
	ContainerId *string `json:"container_id,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// pod id
	PodId *string `json:"pod_id,omitempty"`

	// pod名称
	PodName *string `json:"pod_name,omitempty"`

	// 私有IP地址
	PrivateIp *string `json:"private_ip,omitempty"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 启动命令
	EntryPoint *string `json:"entry_point,omitempty"`
}

func (o LocalImageContainerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalImageContainerInfo struct{}"
	}

	return strings.Join([]string{"LocalImageContainerInfo", string(data)}, " ")
}
