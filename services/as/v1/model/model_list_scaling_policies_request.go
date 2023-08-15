package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScalingPoliciesRequest Request Object
type ListScalingPoliciesRequest struct {

	// 伸缩组ID。
	ScalingGroupId string `json:"scaling_group_id"`

	// 伸缩策略名称。
	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	// 策略类型。
	ScalingPolicyType *ListScalingPoliciesRequestScalingPolicyType `json:"scaling_policy_type,omitempty"`

	// 伸缩策略ID。
	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`

	// 查询的起始行号，默认为0。
	StartNumber *int32 `json:"start_number,omitempty"`

	// 查询记录数，默认20，最大100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScalingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListScalingPoliciesRequest", string(data)}, " ")
}

type ListScalingPoliciesRequestScalingPolicyType struct {
	value string
}

type ListScalingPoliciesRequestScalingPolicyTypeEnum struct {
	ALARM      ListScalingPoliciesRequestScalingPolicyType
	SCHEDULED  ListScalingPoliciesRequestScalingPolicyType
	RECURRENCE ListScalingPoliciesRequestScalingPolicyType
}

func GetListScalingPoliciesRequestScalingPolicyTypeEnum() ListScalingPoliciesRequestScalingPolicyTypeEnum {
	return ListScalingPoliciesRequestScalingPolicyTypeEnum{
		ALARM: ListScalingPoliciesRequestScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: ListScalingPoliciesRequestScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: ListScalingPoliciesRequestScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c ListScalingPoliciesRequestScalingPolicyType) Value() string {
	return c.value
}

func (c ListScalingPoliciesRequestScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScalingPoliciesRequestScalingPolicyType) UnmarshalJSON(b []byte) error {
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
