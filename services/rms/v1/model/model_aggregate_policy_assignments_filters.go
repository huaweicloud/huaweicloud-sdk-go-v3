package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AggregatePolicyAssignmentsFilters 聚合合规规则过滤器。
type AggregatePolicyAssignmentsFilters struct {

	// 源帐号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 合规结果。
	ComplianceState *AggregatePolicyAssignmentsFiltersComplianceState `json:"compliance_state,omitempty"`

	// 合规规则名称
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`
}

func (o AggregatePolicyAssignmentsFilters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregatePolicyAssignmentsFilters struct{}"
	}

	return strings.Join([]string{"AggregatePolicyAssignmentsFilters", string(data)}, " ")
}

type AggregatePolicyAssignmentsFiltersComplianceState struct {
	value string
}

type AggregatePolicyAssignmentsFiltersComplianceStateEnum struct {
	NON_COMPLIANT AggregatePolicyAssignmentsFiltersComplianceState
	COMPLIANT     AggregatePolicyAssignmentsFiltersComplianceState
	NOT_STARTED   AggregatePolicyAssignmentsFiltersComplianceState
}

func GetAggregatePolicyAssignmentsFiltersComplianceStateEnum() AggregatePolicyAssignmentsFiltersComplianceStateEnum {
	return AggregatePolicyAssignmentsFiltersComplianceStateEnum{
		NON_COMPLIANT: AggregatePolicyAssignmentsFiltersComplianceState{
			value: "NonCompliant",
		},
		COMPLIANT: AggregatePolicyAssignmentsFiltersComplianceState{
			value: "Compliant",
		},
		NOT_STARTED: AggregatePolicyAssignmentsFiltersComplianceState{
			value: "NotStarted",
		},
	}
}

func (c AggregatePolicyAssignmentsFiltersComplianceState) Value() string {
	return c.value
}

func (c AggregatePolicyAssignmentsFiltersComplianceState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AggregatePolicyAssignmentsFiltersComplianceState) UnmarshalJSON(b []byte) error {
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
