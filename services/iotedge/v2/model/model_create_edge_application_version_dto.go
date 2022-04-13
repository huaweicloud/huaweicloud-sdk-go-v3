package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type CreateEdgeApplicationVersionDto struct {
	// 应用版本

	Version string `json:"version"`
	// 应用描述

	Description *string `json:"description,omitempty"`
	// 应用部署类型，分为docker容器部署类型和process进程部署类型，兼容之前数据，此字段可以为空，为空情况为docker类型

	DeployType *CreateEdgeApplicationVersionDtoDeployType `json:"deploy_type,omitempty"`

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

func (o CreateEdgeApplicationVersionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeApplicationVersionDto struct{}"
	}

	return strings.Join([]string{"CreateEdgeApplicationVersionDto", string(data)}, " ")
}

type CreateEdgeApplicationVersionDtoDeployType struct {
	value string
}

type CreateEdgeApplicationVersionDtoDeployTypeEnum struct {
	DOCKER  CreateEdgeApplicationVersionDtoDeployType
	PROCESS CreateEdgeApplicationVersionDtoDeployType
}

func GetCreateEdgeApplicationVersionDtoDeployTypeEnum() CreateEdgeApplicationVersionDtoDeployTypeEnum {
	return CreateEdgeApplicationVersionDtoDeployTypeEnum{
		DOCKER: CreateEdgeApplicationVersionDtoDeployType{
			value: "docker",
		},
		PROCESS: CreateEdgeApplicationVersionDtoDeployType{
			value: "process",
		},
	}
}

func (c CreateEdgeApplicationVersionDtoDeployType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEdgeApplicationVersionDtoDeployType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
