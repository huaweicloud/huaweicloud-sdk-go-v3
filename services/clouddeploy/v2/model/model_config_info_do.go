package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 部署参数类
type ConfigInfoDo struct {
	// 部署参数名称，用户可自定义

	Name *string `json:"name,omitempty"`
	// 类型，如果填写name字段，则type必选

	Type *ConfigInfoDoType `json:"type,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 部署参数值，如果填写name字段，则value必选

	Value *string `json:"value,omitempty"`
	// 部署任务id，创建部署任务后由系统自动生成

	TaskId *string `json:"task_id,omitempty"`
	// 表示是否为静态参数，值为1时不支持执行时变更参数，值为0时支持，并且也会把该参数上报流水线

	StaticStatus *ConfigInfoDoStaticStatus `json:"static_status,omitempty"`
	// 当参数类型为enum枚举类型时，必须填写可选值

	Limits *[]ParamTypeLimits `json:"limits,omitempty"`
}

func (o ConfigInfoDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigInfoDo struct{}"
	}

	return strings.Join([]string{"ConfigInfoDo", string(data)}, " ")
}

type ConfigInfoDoType struct {
	value string
}

type ConfigInfoDoTypeEnum struct {
	TEXT       ConfigInfoDoType
	HOST_GROUP ConfigInfoDoType
	ENUM       ConfigInfoDoType
}

func GetConfigInfoDoTypeEnum() ConfigInfoDoTypeEnum {
	return ConfigInfoDoTypeEnum{
		TEXT: ConfigInfoDoType{
			value: "text",
		},
		HOST_GROUP: ConfigInfoDoType{
			value: "host_group",
		},
		ENUM: ConfigInfoDoType{
			value: "enum",
		},
	}
}

func (c ConfigInfoDoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigInfoDoType) UnmarshalJSON(b []byte) error {
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

type ConfigInfoDoStaticStatus struct {
	value int32
}

type ConfigInfoDoStaticStatusEnum struct {
	E_0 ConfigInfoDoStaticStatus
	E_1 ConfigInfoDoStaticStatus
}

func GetConfigInfoDoStaticStatusEnum() ConfigInfoDoStaticStatusEnum {
	return ConfigInfoDoStaticStatusEnum{
		E_0: ConfigInfoDoStaticStatus{
			value: 0,
		}, E_1: ConfigInfoDoStaticStatus{
			value: 1,
		},
	}
}

func (c ConfigInfoDoStaticStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigInfoDoStaticStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
