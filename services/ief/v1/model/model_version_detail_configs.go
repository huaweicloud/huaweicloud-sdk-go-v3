package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 容器特殊参数
type VersionDetailConfigs struct {
	// 默认为false，表示是否开启特权模式

	Privileged *bool `json:"privileged,omitempty"`
	// 默认为true，其中true表示主机网络，而false表示端口映射

	HostNetwork *bool `json:"host_network,omitempty"`
	// 应用实例重启模式： 1. Always：当容器终止退出后，总是重启容器； 2. Onfailure：容器异常退出（退出码非0）时才重启容器； 3. Never：容器终止退出后，不重启容器；

	RestartPolicy *string `json:"restart_policy,omitempty"`
	// 容器端口映射值

	Ports *[]AppVersionPorts `json:"ports,omitempty"`
}

func (o VersionDetailConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionDetailConfigs struct{}"
	}

	return strings.Join([]string{"VersionDetailConfigs", string(data)}, " ")
}
