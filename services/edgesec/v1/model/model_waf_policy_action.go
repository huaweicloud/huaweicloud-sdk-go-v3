package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WafPolicyAction 防护动作
type WafPolicyAction struct {

	// web基础防护动作（log为仅记录、block为拦截）
	Category *WafPolicyActionCategory `json:"category,omitempty"`

	// 攻击惩罚规则ID
	FollowedActionId *string `json:"followed_action_id,omitempty"`
}

func (o WafPolicyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafPolicyAction struct{}"
	}

	return strings.Join([]string{"WafPolicyAction", string(data)}, " ")
}

type WafPolicyActionCategory struct {
	value string
}

type WafPolicyActionCategoryEnum struct {
	BLOCK WafPolicyActionCategory
	LOG   WafPolicyActionCategory
}

func GetWafPolicyActionCategoryEnum() WafPolicyActionCategoryEnum {
	return WafPolicyActionCategoryEnum{
		BLOCK: WafPolicyActionCategory{
			value: "block",
		},
		LOG: WafPolicyActionCategory{
			value: "log",
		},
	}
}

func (c WafPolicyActionCategory) Value() string {
	return c.value
}

func (c WafPolicyActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WafPolicyActionCategory) UnmarshalJSON(b []byte) error {
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
