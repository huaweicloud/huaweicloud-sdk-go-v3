package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyAssignmentRequestBody 规则请求体
type PolicyAssignmentRequestBody struct {

	// 规则类型，包括预定义合规规则(builtin)和用户自定义合规规则(custom)
	PolicyAssignmentType *PolicyAssignmentRequestBodyPolicyAssignmentType `json:"policy_assignment_type,omitempty"`

	// 规则名字
	Name string `json:"name"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 触发周期值，可选值：One_Hour, Three_Hours, Six_Hours, Twelve_Hours, TwentyFour_Hours
	Period *PolicyAssignmentRequestBodyPeriod `json:"period,omitempty"`

	PolicyFilter *PolicyFilterDefinition `json:"policy_filter,omitempty"`

	CustomPolicy *CustomPolicy `json:"custom_policy,omitempty"`

	// 策略定义ID
	PolicyDefinitionId *string `json:"policy_definition_id,omitempty"`

	// 规则参数
	Parameters map[string]PolicyParameterValue `json:"parameters,omitempty"`

	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o PolicyAssignmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAssignmentRequestBody struct{}"
	}

	return strings.Join([]string{"PolicyAssignmentRequestBody", string(data)}, " ")
}

type PolicyAssignmentRequestBodyPolicyAssignmentType struct {
	value string
}

type PolicyAssignmentRequestBodyPolicyAssignmentTypeEnum struct {
	BUILTIN PolicyAssignmentRequestBodyPolicyAssignmentType
	CUSTOM  PolicyAssignmentRequestBodyPolicyAssignmentType
}

func GetPolicyAssignmentRequestBodyPolicyAssignmentTypeEnum() PolicyAssignmentRequestBodyPolicyAssignmentTypeEnum {
	return PolicyAssignmentRequestBodyPolicyAssignmentTypeEnum{
		BUILTIN: PolicyAssignmentRequestBodyPolicyAssignmentType{
			value: "builtin",
		},
		CUSTOM: PolicyAssignmentRequestBodyPolicyAssignmentType{
			value: "custom",
		},
	}
}

func (c PolicyAssignmentRequestBodyPolicyAssignmentType) Value() string {
	return c.value
}

func (c PolicyAssignmentRequestBodyPolicyAssignmentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyAssignmentRequestBodyPolicyAssignmentType) UnmarshalJSON(b []byte) error {
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

type PolicyAssignmentRequestBodyPeriod struct {
	value string
}

type PolicyAssignmentRequestBodyPeriodEnum struct {
	ONE_HOUR          PolicyAssignmentRequestBodyPeriod
	THREE_HOURS       PolicyAssignmentRequestBodyPeriod
	SIX_HOURS         PolicyAssignmentRequestBodyPeriod
	TWELVE_HOURS      PolicyAssignmentRequestBodyPeriod
	TWENTY_FOUR_HOURS PolicyAssignmentRequestBodyPeriod
}

func GetPolicyAssignmentRequestBodyPeriodEnum() PolicyAssignmentRequestBodyPeriodEnum {
	return PolicyAssignmentRequestBodyPeriodEnum{
		ONE_HOUR: PolicyAssignmentRequestBodyPeriod{
			value: "One_Hour",
		},
		THREE_HOURS: PolicyAssignmentRequestBodyPeriod{
			value: "Three_Hours",
		},
		SIX_HOURS: PolicyAssignmentRequestBodyPeriod{
			value: "Six_Hours",
		},
		TWELVE_HOURS: PolicyAssignmentRequestBodyPeriod{
			value: "Twelve_Hours",
		},
		TWENTY_FOUR_HOURS: PolicyAssignmentRequestBodyPeriod{
			value: "TwentyFour_Hours",
		},
	}
}

func (c PolicyAssignmentRequestBodyPeriod) Value() string {
	return c.value
}

func (c PolicyAssignmentRequestBodyPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyAssignmentRequestBodyPeriod) UnmarshalJSON(b []byte) error {
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
