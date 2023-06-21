package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEdgeAppVersionDto struct {

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 部署类型docker|process
	DeployType *string `json:"deploy_type,omitempty"`

	// 是否允许部署多实例
	DeployMultiInstance *bool `json:"deploy_multi_instance,omitempty"`

	ContainerSettings *ContainerSettingsDto `json:"container_settings,omitempty"`

	LivenessProbe *ProbeDto `json:"liveness_probe,omitempty"`

	ReadinessProbe *ProbeDto `json:"readiness_probe,omitempty"`

	// 应用集成的边缘SDK版本
	SdkVersion *string `json:"sdk_version,omitempty"`

	// 架构
	Arch *interface{} `json:"arch,omitempty"`

	// 启动命令
	Command *interface{} `json:"command,omitempty"`

	// 启动参数
	Args *interface{} `json:"args,omitempty"`

	// 应用输出路由端点
	Outputs *interface{} `json:"outputs,omitempty"`

	// 应用输入路由
	Inputs *interface{} `json:"inputs,omitempty"`

	// 应用实现的服务列表
	Services *interface{} `json:"services,omitempty"`

	// 模板id
	TplId *string `json:"tpl_id,omitempty"`
}

func (o UpdateEdgeAppVersionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeAppVersionDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeAppVersionDto", string(data)}, " ")
}
