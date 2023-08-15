package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyStateRequestBody 合规评估结果
type PolicyStateRequestBody struct {
	PolicyResource *PolicyResource `json:"policy_resource"`

	// 触发器类型
	TriggerType PolicyStateRequestBodyTriggerType `json:"trigger_type"`

	// 合规状态
	ComplianceState PolicyStateRequestBodyComplianceState `json:"compliance_state"`

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`

	// 规则名称
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`

	// 合规状态评估时间
	EvaluationTime string `json:"evaluation_time"`

	// 评估校验码
	EvaluationHash string `json:"evaluation_hash"`
}

func (o PolicyStateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyStateRequestBody struct{}"
	}

	return strings.Join([]string{"PolicyStateRequestBody", string(data)}, " ")
}

type PolicyStateRequestBodyTriggerType struct {
	value string
}

type PolicyStateRequestBodyTriggerTypeEnum struct {
	RESOURCE PolicyStateRequestBodyTriggerType
	PERIOD   PolicyStateRequestBodyTriggerType
}

func GetPolicyStateRequestBodyTriggerTypeEnum() PolicyStateRequestBodyTriggerTypeEnum {
	return PolicyStateRequestBodyTriggerTypeEnum{
		RESOURCE: PolicyStateRequestBodyTriggerType{
			value: "resource",
		},
		PERIOD: PolicyStateRequestBodyTriggerType{
			value: "period",
		},
	}
}

func (c PolicyStateRequestBodyTriggerType) Value() string {
	return c.value
}

func (c PolicyStateRequestBodyTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyStateRequestBodyTriggerType) UnmarshalJSON(b []byte) error {
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

type PolicyStateRequestBodyComplianceState struct {
	value string
}

type PolicyStateRequestBodyComplianceStateEnum struct {
	NON_COMPLIANT PolicyStateRequestBodyComplianceState
	COMPLIANT     PolicyStateRequestBodyComplianceState
}

func GetPolicyStateRequestBodyComplianceStateEnum() PolicyStateRequestBodyComplianceStateEnum {
	return PolicyStateRequestBodyComplianceStateEnum{
		NON_COMPLIANT: PolicyStateRequestBodyComplianceState{
			value: "NonCompliant",
		},
		COMPLIANT: PolicyStateRequestBodyComplianceState{
			value: "Compliant",
		},
	}
}

func (c PolicyStateRequestBodyComplianceState) Value() string {
	return c.value
}

func (c PolicyStateRequestBodyComplianceState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyStateRequestBodyComplianceState) UnmarshalJSON(b []byte) error {
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
