package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEdgeAppVersionDto struct {

	// 应用描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 部署类型docker|process
	DeployType *string `json:"deploy_type,omitempty" xml:"deploy_type"`

	// 是否允许部署多实例
	DeployMultiInstance *bool `json:"deploy_multi_instance,omitempty" xml:"deploy_multi_instance"`

	ContainerSettings *ContainerSettingsDto `json:"container_settings,omitempty" xml:"container_settings"`

	LivenessProbe *ProbeDto `json:"liveness_probe,omitempty" xml:"liveness_probe"`

	ReadinessProbe *ProbeDto `json:"readiness_probe,omitempty" xml:"readiness_probe"`

	// 应用集成的边缘升得快版本
	SdkVersion *string `json:"sdk_version,omitempty" xml:"sdk_version"`

	// 架构
	Arch *interface{} `json:"arch,omitempty" xml:"arch"`

	// 启动命令
	Command *interface{} `json:"command,omitempty" xml:"command"`

	// 启动参数
	Args *interface{} `json:"args,omitempty" xml:"args"`

	// 应用输出路由端点
	Outputs *interface{} `json:"outputs,omitempty" xml:"outputs"`

	// 应用输入路由
	Inputs *interface{} `json:"inputs,omitempty" xml:"inputs"`

	// 应用实现的服务列表
	Services *interface{} `json:"services,omitempty" xml:"services"`
}

func (o UpdateEdgeAppVersionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeAppVersionDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeAppVersionDto", string(data)}, " ")
}
