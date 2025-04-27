package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAutoExpandPolicyResponse Response Object
type ShowAutoExpandPolicyResponse struct {

	// **参数解释**：  自动扩容策略开关。
	SwitchOption *bool `json:"switch_option,omitempty"`

	// **参数解释**：  存储自动扩容上限，单位GB。
	LimitSize *int32 `json:"limit_size,omitempty"`

	// **参数解释**：  可用存储空间率。
	TriggerAvailablePercent *ShowAutoExpandPolicyResponseTriggerAvailablePercent `json:"trigger_available_percent,omitempty"`

	// **参数解释**：  扩容步长百分比。
	StepPercent    *int32 `json:"step_percent,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAutoExpandPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoExpandPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoExpandPolicyResponse", string(data)}, " ")
}

type ShowAutoExpandPolicyResponseTriggerAvailablePercent struct {
	value int32
}

type ShowAutoExpandPolicyResponseTriggerAvailablePercentEnum struct {
	E_10 ShowAutoExpandPolicyResponseTriggerAvailablePercent
	E_15 ShowAutoExpandPolicyResponseTriggerAvailablePercent
	E_20 ShowAutoExpandPolicyResponseTriggerAvailablePercent
}

func GetShowAutoExpandPolicyResponseTriggerAvailablePercentEnum() ShowAutoExpandPolicyResponseTriggerAvailablePercentEnum {
	return ShowAutoExpandPolicyResponseTriggerAvailablePercentEnum{
		E_10: ShowAutoExpandPolicyResponseTriggerAvailablePercent{
			value: 10,
		}, E_15: ShowAutoExpandPolicyResponseTriggerAvailablePercent{
			value: 15,
		}, E_20: ShowAutoExpandPolicyResponseTriggerAvailablePercent{
			value: 20,
		},
	}
}

func (c ShowAutoExpandPolicyResponseTriggerAvailablePercent) Value() int32 {
	return c.value
}

func (c ShowAutoExpandPolicyResponseTriggerAvailablePercent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAutoExpandPolicyResponseTriggerAvailablePercent) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
