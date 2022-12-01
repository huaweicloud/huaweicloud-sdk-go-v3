package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostContainerPortMapping struct {

	// 容器端口,1到65535之间的整数
	ContainerPort int64 `json:"container_port"`

	// 对应主机网卡地址，形如192.168.0.1
	HostIp *string `json:"host_ip,omitempty"`

	// 主机端口,1到65535之间的整数，与主机端口范围二选一
	HostPort *int64 `json:"host_port,omitempty"`

	HostPortRange *HostPortRange `json:"host_port_range,omitempty"`
}

func (o HostContainerPortMapping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostContainerPortMapping struct{}"
	}

	return strings.Join([]string{"HostContainerPortMapping", string(data)}, " ")
}
