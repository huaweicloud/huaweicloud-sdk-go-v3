package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateEdgeApplicationVersionStateResponse struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// 应用名称
	Version *string `json:"version,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 部署类型docker|process
	DeployType *string `json:"deploy_type,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 应用版本状态
	State *UpdateEdgeApplicationVersionStateResponseState `json:"state,omitempty"`

	// 架构
	Arch *[]string `json:"arch,omitempty"`

	// 发布时间
	PublishTime *string `json:"publish_time,omitempty"`

	// 下线时间
	OffShelfTime   *string `json:"off_shelf_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEdgeApplicationVersionStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeApplicationVersionStateResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeApplicationVersionStateResponse", string(data)}, " ")
}

type UpdateEdgeApplicationVersionStateResponseState struct {
	value string
}

type UpdateEdgeApplicationVersionStateResponseStateEnum struct {
	DRAFT     UpdateEdgeApplicationVersionStateResponseState
	PUBLISHED UpdateEdgeApplicationVersionStateResponseState
	OFF_SHELF UpdateEdgeApplicationVersionStateResponseState
}

func GetUpdateEdgeApplicationVersionStateResponseStateEnum() UpdateEdgeApplicationVersionStateResponseStateEnum {
	return UpdateEdgeApplicationVersionStateResponseStateEnum{
		DRAFT: UpdateEdgeApplicationVersionStateResponseState{
			value: "DRAFT",
		},
		PUBLISHED: UpdateEdgeApplicationVersionStateResponseState{
			value: "PUBLISHED",
		},
		OFF_SHELF: UpdateEdgeApplicationVersionStateResponseState{
			value: "OFF_SHELF",
		},
	}
}

func (c UpdateEdgeApplicationVersionStateResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEdgeApplicationVersionStateResponseState) UnmarshalJSON(b []byte) error {
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
