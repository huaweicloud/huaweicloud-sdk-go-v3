package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerPortInfo 容器端口信息
type ContainerPortInfo struct {

	// 容器端口
	ContainerPort *int32 `json:"container_port,omitempty"`

	// 主机IP
	HostIp *string `json:"host_ip,omitempty"`

	// 主机端口
	HostPort *int32 `json:"host_port,omitempty"`

	// 端口名称
	PortName *string `json:"port_name,omitempty"`

	// 端口协议，举例如下 -TCP,默认TCP -UDP
	Protocol *string `json:"protocol,omitempty"`
}

func (o ContainerPortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerPortInfo struct{}"
	}

	return strings.Join([]string{"ContainerPortInfo", string(data)}, " ")
}
