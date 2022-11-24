package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用模板网络参数配置
type AppConfigs struct {

	// 默认为false，表示是否开启特权模式
	Privileged *bool `json:"privileged,omitempty"`

	// 默认为true，其中true表示主机网络，而false表示端口映射
	HostNetwork *bool `json:"host_network,omitempty"`

	// 应用实例重启模式： - Always：当容器终止退出后，总是重启容器 - Onfailure：容器异常退出（退出码非0）时才重启容器 - Never：容器终止退出后，不重启容器
	RestartPolicy *string `json:"restart_policy,omitempty"`

	// 容器端口映射值
	Ports *[]Ports `json:"ports,omitempty"`

	// 应用实例是否与主机共PID命名空间，默认值false
	HostPid *string `json:"host_pid,omitempty"`
}

func (o AppConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppConfigs struct{}"
	}

	return strings.Join([]string{"AppConfigs", string(data)}, " ")
}
