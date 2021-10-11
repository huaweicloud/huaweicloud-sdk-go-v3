package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowEdgeApplicationVersionResponse struct {
	// 应用ID

	EdgeAppId *string `json:"edge_app_id,omitempty"`
	// 应用名称

	Name *string `json:"name,omitempty"`
	// 部署类型docker|process

	DeployType *string `json:"deploy_type,omitempty"`
	// 应用版本

	Version *string `json:"version,omitempty"`
	// 应用描述

	Description *string `json:"description,omitempty"`
	// 创建时间

	CreateTime *string `json:"create_time,omitempty"`
	// 最后一次修改时间

	UpdateTime *string `json:"update_time,omitempty"`
	// 应用版本状态

	State *ShowEdgeApplicationVersionResponseState `json:"state,omitempty"`

	LivenessProbe *ProbeDto `json:"liveness_probe,omitempty"`

	ReadinessProbe *ProbeDto `json:"readiness_probe,omitempty"`
	// 架构

	Arch *[]string `json:"arch,omitempty"`
	// 启动命令

	Command *[]string `json:"command,omitempty"`
	// 启动参数

	Args *[]string `json:"args,omitempty"`

	ContainerSettings *ContainerSettingsDto `json:"container_settings,omitempty"`
	// 应用输出路由端点

	Outputs *[]string `json:"outputs,omitempty"`
	// 应用输入路由

	Inputs *[]string `json:"inputs,omitempty"`
	// 应用实现的服务列表

	Services *[]string `json:"services,omitempty"`
	// 发布时间

	PublishTime *string `json:"publish_time,omitempty"`
	// 下线时间

	OffShelfTime   *string `json:"off_shelf_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEdgeApplicationVersionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEdgeApplicationVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeApplicationVersionResponse", string(data)}, " ")
}

type ShowEdgeApplicationVersionResponseState struct {
	value string
}

type ShowEdgeApplicationVersionResponseStateEnum struct {
	DRAFT     ShowEdgeApplicationVersionResponseState
	PUBLISHED ShowEdgeApplicationVersionResponseState
	OFF_SHELF ShowEdgeApplicationVersionResponseState
}

func GetShowEdgeApplicationVersionResponseStateEnum() ShowEdgeApplicationVersionResponseStateEnum {
	return ShowEdgeApplicationVersionResponseStateEnum{
		DRAFT: ShowEdgeApplicationVersionResponseState{
			value: "DRAFT",
		},
		PUBLISHED: ShowEdgeApplicationVersionResponseState{
			value: "PUBLISHED",
		},
		OFF_SHELF: ShowEdgeApplicationVersionResponseState{
			value: "OFF_SHELF",
		},
	}
}

func (c ShowEdgeApplicationVersionResponseState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowEdgeApplicationVersionResponseState) UnmarshalJSON(b []byte) error {
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
