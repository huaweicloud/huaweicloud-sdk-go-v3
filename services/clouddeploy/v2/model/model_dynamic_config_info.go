package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 部署任务执行时传递的参数
type DynamicConfigInfo struct {
	// 执行部署任务时传递的参数名称

	Name *string `json:"name,omitempty"`
	// 执行部署任务时传递的参数值

	Value *string `json:"value,omitempty"`
	// 参数值为枚举类型时，返回可选值列表

	Limits *[]ParamTypeLimits `json:"limits,omitempty"`
	// 类型，如果填写动态参数，则类型必选

	Type *DynamicConfigInfoType `json:"type,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
}

func (o DynamicConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DynamicConfigInfo struct{}"
	}

	return strings.Join([]string{"DynamicConfigInfo", string(data)}, " ")
}

type DynamicConfigInfoType struct {
	value string
}

type DynamicConfigInfoTypeEnum struct {
	TEXT       DynamicConfigInfoType
	HOST_GROUP DynamicConfigInfoType
	ENCRYPT    DynamicConfigInfoType
}

func GetDynamicConfigInfoTypeEnum() DynamicConfigInfoTypeEnum {
	return DynamicConfigInfoTypeEnum{
		TEXT: DynamicConfigInfoType{
			value: "text",
		},
		HOST_GROUP: DynamicConfigInfoType{
			value: "host_group",
		},
		ENCRYPT: DynamicConfigInfoType{
			value: "encrypt",
		},
	}
}

func (c DynamicConfigInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DynamicConfigInfoType) UnmarshalJSON(b []byte) error {
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
