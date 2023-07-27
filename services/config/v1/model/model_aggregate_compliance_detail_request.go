package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AggregateComplianceDetailRequest 查询指定聚合合规规则评估结果详情请求体。
type AggregateComplianceDetailRequest struct {

	// 资源聚合器ID
	AggregatorId string `json:"aggregator_id"`

	// 源帐号ID
	AccountId *string `json:"account_id,omitempty"`

	// 合规结果。
	ComplianceState *AggregateComplianceDetailRequestComplianceState `json:"compliance_state,omitempty"`

	// 合规规则名称
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o AggregateComplianceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregateComplianceDetailRequest struct{}"
	}

	return strings.Join([]string{"AggregateComplianceDetailRequest", string(data)}, " ")
}

type AggregateComplianceDetailRequestComplianceState struct {
	value string
}

type AggregateComplianceDetailRequestComplianceStateEnum struct {
	NON_COMPLIANT AggregateComplianceDetailRequestComplianceState
	COMPLIANT     AggregateComplianceDetailRequestComplianceState
}

func GetAggregateComplianceDetailRequestComplianceStateEnum() AggregateComplianceDetailRequestComplianceStateEnum {
	return AggregateComplianceDetailRequestComplianceStateEnum{
		NON_COMPLIANT: AggregateComplianceDetailRequestComplianceState{
			value: "NonCompliant",
		},
		COMPLIANT: AggregateComplianceDetailRequestComplianceState{
			value: "Compliant",
		},
	}
}

func (c AggregateComplianceDetailRequestComplianceState) Value() string {
	return c.value
}

func (c AggregateComplianceDetailRequestComplianceState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AggregateComplianceDetailRequestComplianceState) UnmarshalJSON(b []byte) error {
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
