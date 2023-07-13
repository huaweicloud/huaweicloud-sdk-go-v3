package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ManagedPolicyAssignmentMetadata 托管规则元数据。
type ManagedPolicyAssignmentMetadata struct {

	// 规则描述。
	Description *string `json:"description,omitempty"`

	// 触发周期。
	Period *ManagedPolicyAssignmentMetadataPeriod `json:"period,omitempty"`

	// 输入参数。
	Parameters map[string]PolicyParameterValue `json:"parameters,omitempty"`

	PolicyFilter *PolicyFilterDefinition `json:"policy_filter,omitempty"`

	// 预定义策略标识符。
	PolicyDefinitionId string `json:"policy_definition_id"`
}

func (o ManagedPolicyAssignmentMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagedPolicyAssignmentMetadata struct{}"
	}

	return strings.Join([]string{"ManagedPolicyAssignmentMetadata", string(data)}, " ")
}

type ManagedPolicyAssignmentMetadataPeriod struct {
	value string
}

type ManagedPolicyAssignmentMetadataPeriodEnum struct {
	ONE_HOUR          ManagedPolicyAssignmentMetadataPeriod
	THREE_HOURS       ManagedPolicyAssignmentMetadataPeriod
	SIX_HOURS         ManagedPolicyAssignmentMetadataPeriod
	TWELVE_HOURS      ManagedPolicyAssignmentMetadataPeriod
	TWENTY_FOUR_HOURS ManagedPolicyAssignmentMetadataPeriod
}

func GetManagedPolicyAssignmentMetadataPeriodEnum() ManagedPolicyAssignmentMetadataPeriodEnum {
	return ManagedPolicyAssignmentMetadataPeriodEnum{
		ONE_HOUR: ManagedPolicyAssignmentMetadataPeriod{
			value: "One_Hour",
		},
		THREE_HOURS: ManagedPolicyAssignmentMetadataPeriod{
			value: "Three_Hours",
		},
		SIX_HOURS: ManagedPolicyAssignmentMetadataPeriod{
			value: "Six_Hours",
		},
		TWELVE_HOURS: ManagedPolicyAssignmentMetadataPeriod{
			value: "Twelve_Hours",
		},
		TWENTY_FOUR_HOURS: ManagedPolicyAssignmentMetadataPeriod{
			value: "TwentyFour_Hours",
		},
	}
}

func (c ManagedPolicyAssignmentMetadataPeriod) Value() string {
	return c.value
}

func (c ManagedPolicyAssignmentMetadataPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ManagedPolicyAssignmentMetadataPeriod) UnmarshalJSON(b []byte) error {
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
