package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowEffectivePoliciesRequest struct {

	// 帐号的唯一标识符（ID）。当前还不支持指定根、组织单元。
	EntityId string `json:"entity_id"`

	// 策略类型的名称，tag_policy标签策略。
	PolicyType ShowEffectivePoliciesRequestPolicyType `json:"policy_type"`
}

func (o ShowEffectivePoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEffectivePoliciesRequest struct{}"
	}

	return strings.Join([]string{"ShowEffectivePoliciesRequest", string(data)}, " ")
}

type ShowEffectivePoliciesRequestPolicyType struct {
	value string
}

type ShowEffectivePoliciesRequestPolicyTypeEnum struct {
	TAG_POLICY ShowEffectivePoliciesRequestPolicyType
}

func GetShowEffectivePoliciesRequestPolicyTypeEnum() ShowEffectivePoliciesRequestPolicyTypeEnum {
	return ShowEffectivePoliciesRequestPolicyTypeEnum{
		TAG_POLICY: ShowEffectivePoliciesRequestPolicyType{
			value: "tag_policy",
		},
	}
}

func (c ShowEffectivePoliciesRequestPolicyType) Value() string {
	return c.value
}

func (c ShowEffectivePoliciesRequestPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEffectivePoliciesRequestPolicyType) UnmarshalJSON(b []byte) error {
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
