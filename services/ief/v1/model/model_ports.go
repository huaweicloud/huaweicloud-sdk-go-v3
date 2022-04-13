package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 容器端口映射值
type Ports struct {
	// 构成一堆映射的容器端口

	ContainerPort *int32 `json:"container_port,omitempty"`
	// 构成一对映射的物理机对应网卡端口

	HostPort *int32 `json:"host_port,omitempty"`
	// 对应网卡地址

	HostIp *string `json:"host_ip,omitempty"`
}

func (o Ports) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ports struct{}"
	}

	return strings.Join([]string{"Ports", string(data)}, " ")
}
