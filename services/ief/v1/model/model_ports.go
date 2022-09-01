package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 容器端口映射值
type Ports struct {

	// 构成一堆映射的容器端口
	ContainerPort *int32 `json:"container_port,omitempty" xml:"container_port"`

	// 构成一对映射的物理机对应网卡端口
	HostPort *int32 `json:"host_port,omitempty" xml:"host_port"`

	// 对应网卡地址
	HostIp *string `json:"host_ip,omitempty" xml:"host_ip"`
}

func (o Ports) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ports struct{}"
	}

	return strings.Join([]string{"Ports", string(data)}, " ")
}
