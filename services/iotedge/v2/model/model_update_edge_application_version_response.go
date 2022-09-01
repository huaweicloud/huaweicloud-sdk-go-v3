package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateEdgeApplicationVersionResponse struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty" xml:"edge_app_id"`

	// 应用名称
	Version *string `json:"version,omitempty" xml:"version"`

	// 应用集成的边缘升得快版本
	SdkVersion *string `json:"sdk_version,omitempty" xml:"sdk_version"`

	// 应用描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 部署类型docker|process
	DeployType *string `json:"deploy_type,omitempty" xml:"deploy_type"`

	// 是否允许部署多实例
	DeployMultiInstance *bool `json:"deploy_multi_instance,omitempty" xml:"deploy_multi_instance"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// 应用版本状态
	State *UpdateEdgeApplicationVersionResponseState `json:"state,omitempty" xml:"state"`

	// 架构
	Arch *[]string `json:"arch,omitempty" xml:"arch"`

	// 发布时间
	PublishTime *string `json:"publish_time,omitempty" xml:"publish_time"`

	// 下线时间
	OffShelfTime   *string `json:"off_shelf_time,omitempty" xml:"off_shelf_time"`
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
