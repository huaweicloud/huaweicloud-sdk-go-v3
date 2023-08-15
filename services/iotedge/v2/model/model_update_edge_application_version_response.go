package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateEdgeApplicationVersionResponse Response Object
type UpdateEdgeApplicationVersionResponse struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用名称
	Version *string `json:"version,omitempty"`

	// 应用集成的边缘SDK版本
	SdkVersion *string `json:"sdk_version,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 部署类型docker|process
	DeployType *string `json:"deploy_type,omitempty"`

	// 是否允许部署多实例
	DeployMultiInstance *bool `json:"deploy_multi_instance,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 应用版本状态
	State *UpdateEdgeApplicationVersionResponseState `json:"state,omitempty"`

	// 架构
	Arch *[]string `json:"arch,omitempty"`

	// 发布时间
	PublishTime *string `json:"publish_time,omitempty"`

	// 下线时间
	OffShelfTime   *string `json:"off_shelf_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEdgeApplicationVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeApplicationVersionResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeApplicationVersionResponse", string(data)}, " ")
}

type UpdateEdgeApplicationVersionResponseState struct {
	value string
}

type UpdateEdgeApplicationVersionResponseStateEnum struct {
	DRAFT     UpdateEdgeApplicationVersionResponseState
	PUBLISHED UpdateEdgeApplicationVersionResponseState
	OFF_SHELF UpdateEdgeApplicationVersionResponseState
}

func GetUpdateEdgeApplicationVersionResponseStateEnum() UpdateEdgeApplicationVersionResponseStateEnum {
	return UpdateEdgeApplicationVersionResponseStateEnum{
		DRAFT: UpdateEdgeApplicationVersionResponseState{
			value: "DRAFT",
		},
		PUBLISHED: UpdateEdgeApplicationVersionResponseState{
			value: "PUBLISHED",
		},
		OFF_SHELF: UpdateEdgeApplicationVersionResponseState{
			value: "OFF_SHELF",
		},
	}
}

func (c UpdateEdgeApplicationVersionResponseState) Value() string {
	return c.value
}

func (c UpdateEdgeApplicationVersionResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEdgeApplicationVersionResponseState) UnmarshalJSON(b []byte) error {
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
