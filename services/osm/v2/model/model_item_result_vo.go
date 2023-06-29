package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ItemResultVo struct {

	// 执行状态
	Status *int32 `json:"status,omitempty"`

	// 是否可见
	Visible *ItemResultVoVisible `json:"visible,omitempty"`

	// 检查项的大类Id
	ResourceId *string `json:"resource_id,omitempty"`

	// 检查项的大类名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 检查项ID
	CheckId *string `json:"check_id,omitempty"`

	// 执行状态
	CheckName *string `json:"check_name,omitempty"`

	// 问题等级 0:未执行，3 正常，4 异常；可用于判断检查项执行结果
	ProblemLevel *ItemResultVoProblemLevel `json:"problem_level,omitempty"`
}

func (o ItemResultVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemResultVo struct{}"
	}

	return strings.Join([]string{"ItemResultVo", string(data)}, " ")
}

type ItemResultVoVisible struct {
	value int32
}

type ItemResultVoVisibleEnum struct {
	E_0 ItemResultVoVisible
	E_1 ItemResultVoVisible
}

func GetItemResultVoVisibleEnum() ItemResultVoVisibleEnum {
	return ItemResultVoVisibleEnum{
		E_0: ItemResultVoVisible{
			value: 0,
		}, E_1: ItemResultVoVisible{
			value: 1,
		},
	}
}

func (c ItemResultVoVisible) Value() int32 {
	return c.value
}

func (c ItemResultVoVisible) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ItemResultVoVisible) UnmarshalJSON(b []byte) error {
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

type ItemResultVoProblemLevel struct {
	value int32
}

type ItemResultVoProblemLevelEnum struct {
	E_0 ItemResultVoProblemLevel
	E_3 ItemResultVoProblemLevel
	E_4 ItemResultVoProblemLevel
}

func GetItemResultVoProblemLevelEnum() ItemResultVoProblemLevelEnum {
	return ItemResultVoProblemLevelEnum{
		E_0: ItemResultVoProblemLevel{
			value: 0,
		}, E_3: ItemResultVoProblemLevel{
			value: 3,
		}, E_4: ItemResultVoProblemLevel{
			value: 4,
		},
	}
}

func (c ItemResultVoProblemLevel) Value() int32 {
	return c.value
}

func (c ItemResultVoProblemLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ItemResultVoProblemLevel) UnmarshalJSON(b []byte) error {
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
