package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEffectivePoliciesRequest Request Object
type ShowEffectivePoliciesRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 账号的唯一标识符（ID）。当前还不支持指定根、组织单元。
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
