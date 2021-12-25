package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerResp struct {
	// 容器启动参数，字符总长度最大为65536

	Args *[]string `json:"args,omitempty"`
	// 容器启动命令，字符总长度最大为65536。 command支持使用数组定义多条命令，但在IEF控制台界面只会显示第一条命令。

	Command *[]string `json:"command,omitempty"`
	// 容器名称，只允许英文小写字母、数字、中划线，最大长度32, 英文小写字母或数字开头和结尾

	Name string `json:"name"`
	// 环境变量

	Envs *[]EnvPods `json:"envs,omitempty"`
	// 容器镜像URL

	ImageUrl string `json:"image_url"`
	// 容器镜像版本

	Version *string `json:"version,omitempty"`

	LivenessProbe *Probe `json:"liveness_probe,omitempty"`

	ReadinessProbe *Probe `json:"readiness_probe,omitempty"`
	// 容器端口映射值

	Ports *[]HostContainerPort `json:"ports,omitempty"`

	Resources *DeploymentResources `json:"resources,omitempty"`
	// 卷配置

	Volumes *[]Volumes `json:"volumes,omitempty"`
	// 容器重启次数

	Restarts *int64 `json:"restarts,omitempty"`
	// 容器故障详情

	Message *string `json:"message,omitempty"`
	// 容器故障原因

	Reason *string `json:"reason,omitempty"`
	// 健康检查结果

	IsReady *string `json:"is_ready,omitempty"`
	// 是否启用特权容器,默认值false

	Privileged *bool `json:"privileged,omitempty"`
	// 容器ID

	ContainerId *string `json:"container_id,omitempty"`
	// 容器状态

	State *string `json:"state,omitempty"`
	// npu类型，支持D310类型和D910类型。 - D310表示D310类型。 - D910表示D910类型。 - 不填表示为D310类型。

	NpuType *string `json:"npu_type,omitempty"`
}

func (o ContainerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerResp struct{}"
	}

	return strings.Join([]string{"ContainerResp", string(data)}, " ")
}
