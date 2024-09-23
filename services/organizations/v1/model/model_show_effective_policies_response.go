package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowEffectivePoliciesResponse Response Object
type ShowEffectivePoliciesResponse struct {

	// 有效策略最后更新时间。
	LastUpdatedAt *sdktime.SdkTime `json:"last_updated_at,omitempty"`

	// 有效策略文本内容。
	PolicyContent *string `json:"policy_content,omitempty"`

	// 策略类型的名称。tag_policy标签策略。
	PolicyType *ShowEffectivePoliciesResponsePolicyType `json:"policy_type,omitempty"`

	// 根、组织单元或账号的唯一标识符（ID）。
	EntityId       *string `json:"entity_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEffectivePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEffectivePoliciesResponse struct{}"
	}

	return strings.Join([]string{"ShowEffectivePoliciesResponse", string(data)}, " ")
}

type ShowEffectivePoliciesResponsePolicyType struct {
	value string
}

type ShowEffectivePoliciesResponsePolicyTypeEnum struct {
	TAG_POLICY ShowEffectivePoliciesResponsePolicyType
}

func GetShowEffectivePoliciesResponsePolicyTypeEnum() ShowEffectivePoliciesResponsePolicyTypeEnum {
	return ShowEffectivePoliciesResponsePolicyTypeEnum{
		TAG_POLICY: ShowEffectivePoliciesResponsePolicyType{
			value: "tag_policy",
		},
	}
}

func (c ShowEffectivePoliciesResponsePolicyType) Value() string {
	return c.value
}

func (c ShowEffectivePoliciesResponsePolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEffectivePoliciesResponsePolicyType) UnmarshalJSON(b []byte) error {
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
