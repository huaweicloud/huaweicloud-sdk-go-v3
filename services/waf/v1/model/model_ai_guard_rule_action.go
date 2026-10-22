package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AiGuardRuleAction 规则动作配置
type AiGuardRuleAction struct {

	// 动作分类，可选：log(仅记录), block(拦截),desensitize(脱敏)
	Category *AiGuardRuleActionCategory `json:"category,omitempty"`
}

func (o AiGuardRuleAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiGuardRuleAction struct{}"
	}

	return strings.Join([]string{"AiGuardRuleAction", string(data)}, " ")
}

type AiGuardRuleActionCategory struct {
	value string
}

type AiGuardRuleActionCategoryEnum struct {
	LOG         AiGuardRuleActionCategory
	BLOCK       AiGuardRuleActionCategory
	DESENSITIZE AiGuardRuleActionCategory
}

func GetAiGuardRuleActionCategoryEnum() AiGuardRuleActionCategoryEnum {
	return AiGuardRuleActionCategoryEnum{
		LOG: AiGuardRuleActionCategory{
			value: "log",
		},
		BLOCK: AiGuardRuleActionCategory{
			value: "block",
		},
		DESENSITIZE: AiGuardRuleActionCategory{
			value: "desensitize",
		},
	}
}

func (c AiGuardRuleActionCategory) Value() string {
	return c.value
}

func (c AiGuardRuleActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AiGuardRuleActionCategory) UnmarshalJSON(b []byte) error {
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
