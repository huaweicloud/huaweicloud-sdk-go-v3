package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyAction 防护动作
type PolicyAction struct {

	// web基础防护动作（log为仅记录、block为拦截）
	Category *PolicyActionCategory `json:"category,omitempty"`

	// 攻击惩罚规则ID
	FollowedActionId *string `json:"followed_action_id,omitempty"`
}

func (o PolicyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAction struct{}"
	}

	return strings.Join([]string{"PolicyAction", string(data)}, " ")
}

type PolicyActionCategory struct {
	value string
}

type PolicyActionCategoryEnum struct {
	BLOCK PolicyActionCategory
	LOG   PolicyActionCategory
}

func GetPolicyActionCategoryEnum() PolicyActionCategoryEnum {
	return PolicyActionCategoryEnum{
		BLOCK: PolicyActionCategory{
			value: "block",
		},
		LOG: PolicyActionCategory{
			value: "log",
		},
	}
}

func (c PolicyActionCategory) Value() string {
	return c.value
}

func (c PolicyActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyActionCategory) UnmarshalJSON(b []byte) error {
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
