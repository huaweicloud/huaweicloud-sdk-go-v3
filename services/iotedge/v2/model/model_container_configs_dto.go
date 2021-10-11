package model

import (
	"encoding/json"

	"strings"
)

type ContainerConfigsDto struct {
	// 开启容器特权模式

	Privileged *bool `json:"privileged,omitempty"`
	// 是否使用主机网络模式

	HostNetwork *bool `json:"host_network,omitempty"`
	// 重启策略，容器执行健康检查后失败后的策略

	RestartPolicy string `json:"restart_policy"`
	// 容器端口映射值

	ContainerPortList *[]ContainerPortDto `json:"container_port_list,omitempty"`
}

func (o ContainerConfigsDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ContainerConfigsDto struct{}"
	}

	return strings.Join([]string{"ContainerConfigsDto", string(data)}, " ")
}
