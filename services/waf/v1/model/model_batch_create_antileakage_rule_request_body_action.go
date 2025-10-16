package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateAntileakageRuleRequestBodyAction 规则命中后操作对象
type BatchCreateAntileakageRuleRequestBodyAction struct {

	// 操作类型。   - “block”：过滤。   - “log”：仅记录
	Category BatchCreateAntileakageRuleRequestBodyActionCategory `json:"category"`
}

func (o BatchCreateAntileakageRuleRequestBodyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAntileakageRuleRequestBodyAction struct{}"
	}

	return strings.Join([]string{"BatchCreateAntileakageRuleRequestBodyAction", string(data)}, " ")
}

type BatchCreateAntileakageRuleRequestBodyActionCategory struct {
	value string
}

type BatchCreateAntileakageRuleRequestBodyActionCategoryEnum struct {
	BLOCK BatchCreateAntileakageRuleRequestBodyActionCategory
	LOG   BatchCreateAntileakageRuleRequestBodyActionCategory
}

func GetBatchCreateAntileakageRuleRequestBodyActionCategoryEnum() BatchCreateAntileakageRuleRequestBodyActionCategoryEnum {
	return BatchCreateAntileakageRuleRequestBodyActionCategoryEnum{
		BLOCK: BatchCreateAntileakageRuleRequestBodyActionCategory{
			value: "block",
		},
		LOG: BatchCreateAntileakageRuleRequestBodyActionCategory{
			value: "log",
		},
	}
}

func (c BatchCreateAntileakageRuleRequestBodyActionCategory) Value() string {
	return c.value
}

func (c BatchCreateAntileakageRuleRequestBodyActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateAntileakageRuleRequestBodyActionCategory) UnmarshalJSON(b []byte) error {
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
