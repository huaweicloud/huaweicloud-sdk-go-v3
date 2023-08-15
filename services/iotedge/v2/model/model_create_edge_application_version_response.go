package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateEdgeApplicationVersionResponse Response Object
type CreateEdgeApplicationVersionResponse struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 部署类型docker|process
	DeployType *string `json:"deploy_type,omitempty"`

	// 是否允许部署多实例
	DeployMultiInstance *bool `json:"deploy_multi_instance,omitempty"`

	// 应用版本
	Version *string `json:"version,omitempty"`

	// 应用集成的边缘SDK版本
	SdkVersion *string `json:"sdk_version,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 应用版本状态
	State *CreateEdgeApplicationVersionResponseState `json:"state,omitempty"`

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
	OffShelfTime *string `json:"off_shelf_time,omitempty"`

	// 驱动厂商
	Supplier *string `json:"supplier,omitempty"`

	// 模板id
	TplId          *string `json:"tpl_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEdgeApplicationVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeApplicationVersionResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeApplicationVersionResponse", string(data)}, " ")
}

type CreateEdgeApplicationVersionResponseState struct {
	value string
}

type CreateEdgeApplicationVersionResponseStateEnum struct {
	DRAFT     CreateEdgeApplicationVersionResponseState
	PUBLISHED CreateEdgeApplicationVersionResponseState
	OFF_SHELF CreateEdgeApplicationVersionResponseState
}

func GetCreateEdgeApplicationVersionResponseStateEnum() CreateEdgeApplicationVersionResponseStateEnum {
	return CreateEdgeApplicationVersionResponseStateEnum{
		DRAFT: CreateEdgeApplicationVersionResponseState{
			value: "DRAFT",
		},
		PUBLISHED: CreateEdgeApplicationVersionResponseState{
			value: "PUBLISHED",
		},
		OFF_SHELF: CreateEdgeApplicationVersionResponseState{
			value: "OFF_SHELF",
		},
	}
}

func (c CreateEdgeApplicationVersionResponseState) Value() string {
	return c.value
}

func (c CreateEdgeApplicationVersionResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEdgeApplicationVersionResponseState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
