package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowIaConfigResponse struct {
	// 配置ID

	Id *string `json:"id,omitempty"`
	// 配置项名称

	Name *string `json:"name,omitempty"`
	// 配置项详情

	Value *string `json:"value,omitempty"`
	// 配置项描述

	Description *string `json:"description,omitempty"`
	// 版本号

	Version *int64 `json:"version,omitempty"`
	// 下发状态

	State *ShowIaConfigResponseState `json:"state,omitempty"`
	// 创建时间

	CreateTime *string `json:"create_time,omitempty"`
	// 更新时间

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIaConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIaConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowIaConfigResponse", string(data)}, " ")
}

type ShowIaConfigResponseState struct {
	value string
}

type ShowIaConfigResponseStateEnum struct {
	SUCCESS ShowIaConfigResponseState
	SENDING ShowIaConfigResponseState
}

func GetShowIaConfigResponseStateEnum() ShowIaConfigResponseStateEnum {
	return ShowIaConfigResponseStateEnum{
		SUCCESS: ShowIaConfigResponseState{
			value: "SUCCESS",
		},
		SENDING: ShowIaConfigResponseState{
			value: "SENDING",
		},
	}
}

func (c ShowIaConfigResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowIaConfigResponseState) UnmarshalJSON(b []byte) error {
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
