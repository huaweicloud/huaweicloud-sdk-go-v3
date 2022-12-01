package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PodConfig struct {

	// 应用实例是否启用主机网络，不启用则使用端口映射,默认值false
	HostNetwork *bool `json:"host_network,omitempty"`

	// 应用实例是否与主机共PID命名空间,默认值false
	HostPid *bool `json:"host_pid,omitempty"`

	// 应用实例故障是否迁移,指定节点组部署时必选，默认值false
	Migration *bool `json:"migration,omitempty"`

	// 应用实例重启策略,可选值Always、OnFailure、Never
	RestartPolicy *string `json:"restart_policy,omitempty"`

	// 应用实例故障容忍时间,容忍时间到达后迁移应用实例，只在指定节点组部署时生效
	TolerationSeconds *int32 `json:"toleration_seconds,omitempty"`
}

func (o PodConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodConfig struct{}"
	}

	return strings.Join([]string{"PodConfig", string(data)}, " ")
}
