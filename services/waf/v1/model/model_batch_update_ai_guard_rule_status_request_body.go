package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchUpdateAiGuardRuleStatusRequestBody 批量更新规则状态请求
type BatchUpdateAiGuardRuleStatusRequestBody struct {

	// 规则ID列表
	RuleIds *[]string `json:"rule_ids,omitempty"`

	// 规则状态，可选：enable(启用), disable(禁用)，当disable时，只有rule_ids，而enable时至少有rule_ids或者temlate_ids
	Status BatchUpdateAiGuardRuleStatusRequestBodyStatus `json:"status"`

	// 系统规则的id
	TemplateIds *[]string `json:"template_ids,omitempty"`
}

func (o BatchUpdateAiGuardRuleStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAiGuardRuleStatusRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateAiGuardRuleStatusRequestBody", string(data)}, " ")
}

type BatchUpdateAiGuardRuleStatusRequestBodyStatus struct {
	value string
}

type BatchUpdateAiGuardRuleStatusRequestBodyStatusEnum struct {
	ENABLE  BatchUpdateAiGuardRuleStatusRequestBodyStatus
	DISABLE BatchUpdateAiGuardRuleStatusRequestBodyStatus
}

func GetBatchUpdateAiGuardRuleStatusRequestBodyStatusEnum() BatchUpdateAiGuardRuleStatusRequestBodyStatusEnum {
	return BatchUpdateAiGuardRuleStatusRequestBodyStatusEnum{
		ENABLE: BatchUpdateAiGuardRuleStatusRequestBodyStatus{
			value: "enable",
		},
		DISABLE: BatchUpdateAiGuardRuleStatusRequestBodyStatus{
			value: "disable",
		},
	}
}

func (c BatchUpdateAiGuardRuleStatusRequestBodyStatus) Value() string {
	return c.value
}

func (c BatchUpdateAiGuardRuleStatusRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateAiGuardRuleStatusRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
