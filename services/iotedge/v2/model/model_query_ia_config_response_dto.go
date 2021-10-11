package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QueryIaConfigResponseDto struct {
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

	State *QueryIaConfigResponseDtoState `json:"state,omitempty"`
	// 创建时间

	CreateTime *string `json:"create_time,omitempty"`
	// 更新时间

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o QueryIaConfigResponseDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryIaConfigResponseDto struct{}"
	}

	return strings.Join([]string{"QueryIaConfigResponseDto", string(data)}, " ")
}

type QueryIaConfigResponseDtoState struct {
	value string
}

type QueryIaConfigResponseDtoStateEnum struct {
	SUCCESS QueryIaConfigResponseDtoState
	SENDING QueryIaConfigResponseDtoState
}

func GetQueryIaConfigResponseDtoStateEnum() QueryIaConfigResponseDtoStateEnum {
	return QueryIaConfigResponseDtoStateEnum{
		SUCCESS: QueryIaConfigResponseDtoState{
			value: "SUCCESS",
		},
		SENDING: QueryIaConfigResponseDtoState{
			value: "SENDING",
		},
	}
}

func (c QueryIaConfigResponseDtoState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *QueryIaConfigResponseDtoState) UnmarshalJSON(b []byte) error {
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
