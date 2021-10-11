package model

import (
	"encoding/json"

	"strings"
)

type UpdateEdgeAppVersionDto struct {
	// 应用描述

	Description *string `json:"description,omitempty"`
	// 部署类型docker|process

	DeployType *string `json:"deploy_type,omitempty"`

	ContainerSettings *ContainerSettingsDto `json:"container_settings"`

	LivenessProbe *ProbeDto `json:"liveness_probe,omitempty"`

	ReadinessProbe *ProbeDto `json:"readiness_probe,omitempty"`
	// 架构

	Arch *interface{} `json:"arch"`
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
}

func (o UpdateEdgeAppVersionDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEdgeAppVersionDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeAppVersionDto", string(data)}, " ")
}
