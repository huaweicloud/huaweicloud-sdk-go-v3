package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAntileakageRuleRequestBodyAction 规则命中后操作对象
type CreateAntileakageRuleRequestBodyAction struct {

	// 操作类型。   - “block”：过滤。   - “log”：仅记录
	Category CreateAntileakageRuleRequestBodyActionCategory `json:"category"`
}

func (o CreateAntileakageRuleRequestBodyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntileakageRuleRequestBodyAction struct{}"
	}

	return strings.Join([]string{"CreateAntileakageRuleRequestBodyAction", string(data)}, " ")
}

type CreateAntileakageRuleRequestBodyActionCategory struct {
	value string
}

type CreateAntileakageRuleRequestBodyActionCategoryEnum struct {
	BLOCK CreateAntileakageRuleRequestBodyActionCategory
	LOG   CreateAntileakageRuleRequestBodyActionCategory
}

func GetCreateAntileakageRuleRequestBodyActionCategoryEnum() CreateAntileakageRuleRequestBodyActionCategoryEnum {
	return CreateAntileakageRuleRequestBodyActionCategoryEnum{
		BLOCK: CreateAntileakageRuleRequestBodyActionCategory{
			value: "block",
		},
		LOG: CreateAntileakageRuleRequestBodyActionCategory{
			value: "log",
		},
	}
}

func (c CreateAntileakageRuleRequestBodyActionCategory) Value() string {
	return c.value
}

func (c CreateAntileakageRuleRequestBodyActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAntileakageRuleRequestBodyActionCategory) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
