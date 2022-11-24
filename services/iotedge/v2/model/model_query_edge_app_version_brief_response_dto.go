package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QueryEdgeAppVersionBriefResponseDto struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// 应用名称
	Version *string `json:"version,omitempty"`

	// 应用集成的边缘升得快版本
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
	State *QueryEdgeAppVersionBriefResponseDtoState `json:"state,omitempty"`

	// 架构
	Arch *[]string `json:"arch,omitempty"`

	// 发布时间
	PublishTime *string `json:"publish_time,omitempty"`

	// 下线时间
	OffShelfTime *string `json:"off_shelf_time,omitempty"`
}

func (o QueryEdgeAppVersionBriefResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryEdgeAppVersionBriefResponseDto struct{}"
	}

	return strings.Join([]string{"QueryEdgeAppVersionBriefResponseDto", string(data)}, " ")
}

type QueryEdgeAppVersionBriefResponseDtoState struct {
	value string
}

type QueryEdgeAppVersionBriefResponseDtoStateEnum struct {
	DRAFT     QueryEdgeAppVersionBriefResponseDtoState
	PUBLISHED QueryEdgeAppVersionBriefResponseDtoState
	OFF_SHELF QueryEdgeAppVersionBriefResponseDtoState
}

func GetQueryEdgeAppVersionBriefResponseDtoStateEnum() QueryEdgeAppVersionBriefResponseDtoStateEnum {
	return QueryEdgeAppVersionBriefResponseDtoStateEnum{
		DRAFT: QueryEdgeAppVersionBriefResponseDtoState{
			value: "DRAFT",
		},
		PUBLISHED: QueryEdgeAppVersionBriefResponseDtoState{
			value: "PUBLISHED",
		},
		OFF_SHELF: QueryEdgeAppVersionBriefResponseDtoState{
			value: "OFF_SHELF",
		},
	}
}

func (c QueryEdgeAppVersionBriefResponseDtoState) Value() string {
	return c.value
}

func (c QueryEdgeAppVersionBriefResponseDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryEdgeAppVersionBriefResponseDtoState) UnmarshalJSON(b []byte) error {
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
