package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchUpdateBotMRuleActionRequestBody **参数解释：** 批量更新规则防护动作的请求体，用于统一修改多条BotM规则的防护动作 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type BatchUpdateBotMRuleActionRequestBody struct {

	// **参数解释：** 批量更新的规则ID，包含需要修改防护动作的多条BotM规则ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleIds *[]string `json:"rule_ids,omitempty"`

	// **参数解释：** 批量更新的防护动作对应的数字，101-拦截、102-放行、103-观察、104-验证码、201-重定向、202-返回自定义页面、301-拉黑IP **约束限制：** 不涉及 **取值范围：** - 101：拦截 - 102：放行 - 103：观察 - 104：验证码 - 201：重定向 - 202：返回自定义页面 - 301：拉黑IP **默认取值：** 不涉及
	DefenseAction *BatchUpdateBotMRuleActionRequestBodyDefenseAction `json:"defense_action,omitempty"`
}

func (o BatchUpdateBotMRuleActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateBotMRuleActionRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateBotMRuleActionRequestBody", string(data)}, " ")
}

type BatchUpdateBotMRuleActionRequestBodyDefenseAction struct {
	value int32
}

type BatchUpdateBotMRuleActionRequestBodyDefenseActionEnum struct {
	E_101 BatchUpdateBotMRuleActionRequestBodyDefenseAction
	E_102 BatchUpdateBotMRuleActionRequestBodyDefenseAction
	E_103 BatchUpdateBotMRuleActionRequestBodyDefenseAction
	E_104 BatchUpdateBotMRuleActionRequestBodyDefenseAction
	E_201 BatchUpdateBotMRuleActionRequestBodyDefenseAction
	E_202 BatchUpdateBotMRuleActionRequestBodyDefenseAction
	E_301 BatchUpdateBotMRuleActionRequestBodyDefenseAction
}

func GetBatchUpdateBotMRuleActionRequestBodyDefenseActionEnum() BatchUpdateBotMRuleActionRequestBodyDefenseActionEnum {
	return BatchUpdateBotMRuleActionRequestBodyDefenseActionEnum{
		E_101: BatchUpdateBotMRuleActionRequestBodyDefenseAction{
			value: 101,
		}, E_102: BatchUpdateBotMRuleActionRequestBodyDefenseAction{
			value: 102,
		}, E_103: BatchUpdateBotMRuleActionRequestBodyDefenseAction{
			value: 103,
		}, E_104: BatchUpdateBotMRuleActionRequestBodyDefenseAction{
			value: 104,
		}, E_201: BatchUpdateBotMRuleActionRequestBodyDefenseAction{
			value: 201,
		}, E_202: BatchUpdateBotMRuleActionRequestBodyDefenseAction{
			value: 202,
		}, E_301: BatchUpdateBotMRuleActionRequestBodyDefenseAction{
			value: 301,
		},
	}
}

func (c BatchUpdateBotMRuleActionRequestBodyDefenseAction) Value() int32 {
	return c.value
}

func (c BatchUpdateBotMRuleActionRequestBodyDefenseAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateBotMRuleActionRequestBodyDefenseAction) UnmarshalJSON(b []byte) error {
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
