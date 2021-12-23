package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerPortDto struct {
	// 构成一堆映射的容器端口

	ContainerPort *int32 `json:"container_port,omitempty"`
	// 构成一对映射的物理机对应网卡端口

	HostPort *int32 `json:"host_port,omitempty"`
	// 对应网卡地址

	HostIp *string `json:"host_ip,omitempty"`
}

func (o ContainerPortDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerPortDto struct{}"
	}

	return strings.Join([]string{"ContainerPortDto", string(data)}, " ")
}
