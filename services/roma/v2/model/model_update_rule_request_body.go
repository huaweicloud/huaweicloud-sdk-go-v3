package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateRuleRequestBody struct {

	// 规则名称，支持英文大小写，数字，下划线和中划线,长度1-64
	Name *string `json:"name,omitempty"`

	// 描述，长度0-200
	Description *string `json:"description,omitempty"`

	// 规则状态 0-启用 1-停用，不填写时默认为0-启用
	Status *UpdateRuleRequestBodyStatus `json:"status,omitempty"`

	// 数据解析状态，0-启用 1-停用，不填写时默认为1-禁用
	DataParsingStatus *UpdateRuleRequestBodyDataParsingStatus `json:"data_parsing_status,omitempty"`

	// SQL查询字段
	SqlField *string `json:"sql_field,omitempty"`

	// SQL查询条件
	SqlWhere *string `json:"sql_where,omitempty"`
}

func (o UpdateRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRuleRequestBody", string(data)}, " ")
}

type UpdateRuleRequestBodyStatus struct {
	value int32
}

type UpdateRuleRequestBodyStatusEnum struct {
	E_0 UpdateRuleRequestBodyStatus
	E_1 UpdateRuleRequestBodyStatus
}

func GetUpdateRuleRequestBodyStatusEnum() UpdateRuleRequestBodyStatusEnum {
	return UpdateRuleRequestBodyStatusEnum{
		E_0: UpdateRuleRequestBodyStatus{
			value: 0,
		}, E_1: UpdateRuleRequestBodyStatus{
			value: 1,
		},
	}
}

func (c UpdateRuleRequestBodyStatus) Value() int32 {
	return c.value
}

func (c UpdateRuleRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleRequestBodyStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdateRuleRequestBodyDataParsingStatus struct {
	value int32
}

type UpdateRuleRequestBodyDataParsingStatusEnum struct {
	E_0 UpdateRuleRequestBodyDataParsingStatus
	E_1 UpdateRuleRequestBodyDataParsingStatus
}

func GetUpdateRuleRequestBodyDataParsingStatusEnum() UpdateRuleRequestBodyDataParsingStatusEnum {
	return UpdateRuleRequestBodyDataParsingStatusEnum{
		E_0: UpdateRuleRequestBodyDataParsingStatus{
			value: 0,
		}, E_1: UpdateRuleRequestBodyDataParsingStatus{
			value: 1,
		},
	}
}

func (c UpdateRuleRequestBodyDataParsingStatus) Value() int32 {
	return c.value
}

func (c UpdateRuleRequestBodyDataParsingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleRequestBodyDataParsingStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
