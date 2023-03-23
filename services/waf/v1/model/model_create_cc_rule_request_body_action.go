package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 请求次数限制到达后采取的防护动作
type CreateCcRuleRequestBodyAction struct {

	// 动作类型：  - captcha：人机验证，阻断后用户需要输入正确的验证码，恢复正确的访问页面。  -block：阻断。   - log: 仅记录   - dynamic_block: 上一个限速周期内，请求频率超过“限速频率”将被阻断，那么在下一个限速周期内，请求频率超过“放行频率”将被阻断。注：只有当cc防护规则模式为高级模式时才支持设置dynamic_block防护动作。
	Category CreateCcRuleRequestBodyActionCategory `json:"category"`

	Detail *CreateCcRuleRequestBodyActionDetail `json:"detail,omitempty"`
}

func (o CreateCcRuleRequestBodyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCcRuleRequestBodyAction struct{}"
	}

	return strings.Join([]string{"CreateCcRuleRequestBodyAction", string(data)}, " ")
}

type CreateCcRuleRequestBodyActionCategory struct {
	value string
}

type CreateCcRuleRequestBodyActionCategoryEnum struct {
	CAPTCHA       CreateCcRuleRequestBodyActionCategory
	BLOCK         CreateCcRuleRequestBodyActionCategory
	LOG           CreateCcRuleRequestBodyActionCategory
	DYNAMIC_BLOCK CreateCcRuleRequestBodyActionCategory
}

func GetCreateCcRuleRequestBodyActionCategoryEnum() CreateCcRuleRequestBodyActionCategoryEnum {
	return CreateCcRuleRequestBodyActionCategoryEnum{
		CAPTCHA: CreateCcRuleRequestBodyActionCategory{
			value: "captcha",
		},
		BLOCK: CreateCcRuleRequestBodyActionCategory{
			value: "block",
		},
		LOG: CreateCcRuleRequestBodyActionCategory{
			value: "log",
		},
		DYNAMIC_BLOCK: CreateCcRuleRequestBodyActionCategory{
			value: "dynamic_block",
		},
	}
}

func (c CreateCcRuleRequestBodyActionCategory) Value() string {
	return c.value
}

func (c CreateCcRuleRequestBodyActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCcRuleRequestBodyActionCategory) UnmarshalJSON(b []byte) error {
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
