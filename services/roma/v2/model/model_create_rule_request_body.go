package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type CreateRuleRequestBody struct {
	// 规则名称，支持英文大小写，数字，下划线和中划线,长度1-64

	Name string `json:"name"`
	// 应用ID

	AppId string `json:"app_id"`
	// 描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 规则状态 0-启用 1-停用，不填写时默认为0-启用

	Status *CreateRuleRequestBodyStatus `json:"status,omitempty"`
	// 数据解析状态，0-启用 1-停用，不填写时默认为1-禁用

	DataParsingStatus *CreateRuleRequestBodyDataParsingStatus `json:"data_parsing_status,omitempty"`
}

func (o CreateRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRuleRequestBody", string(data)}, " ")
}

type CreateRuleRequestBodyStatus struct {
	value int32
}

type CreateRuleRequestBodyStatusEnum struct {
	E_0 CreateRuleRequestBodyStatus
	E_1 CreateRuleRequestBodyStatus
}

func GetCreateRuleRequestBodyStatusEnum() CreateRuleRequestBodyStatusEnum {
	return CreateRuleRequestBodyStatusEnum{
		E_0: CreateRuleRequestBodyStatus{
			value: 0,
		}, E_1: CreateRuleRequestBodyStatus{
			value: 1,
		},
	}
}

func (c CreateRuleRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRuleRequestBodyStatus) UnmarshalJSON(b []byte) error {
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

type CreateRuleRequestBodyDataParsingStatus struct {
	value int32
}

type CreateRuleRequestBodyDataParsingStatusEnum struct {
	E_0 CreateRuleRequestBodyDataParsingStatus
	E_1 CreateRuleRequestBodyDataParsingStatus
}

func GetCreateRuleRequestBodyDataParsingStatusEnum() CreateRuleRequestBodyDataParsingStatusEnum {
	return CreateRuleRequestBodyDataParsingStatusEnum{
		E_0: CreateRuleRequestBodyDataParsingStatus{
			value: 0,
		}, E_1: CreateRuleRequestBodyDataParsingStatus{
			value: 1,
		},
	}
}

func (c CreateRuleRequestBodyDataParsingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRuleRequestBodyDataParsingStatus) UnmarshalJSON(b []byte) error {
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
