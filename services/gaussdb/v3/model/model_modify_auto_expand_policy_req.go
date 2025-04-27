package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ModifyAutoExpandPolicyReq struct {

	// **参数解释**：  自动扩容策略开关。  **取值范围**：  - true：表示开启。 - false：表示关闭。
	SwitchOption *bool `json:"switch_option,omitempty"`

	// **参数解释**：  存储自动扩容上限，需要为10的倍数，单位GB。  **取值范围**：  10 - 最大容量上限。  示例：500
	LimitSize *int32 `json:"limit_size,omitempty"`

	// **参数解释**：  可用存储空间率。  **取值范围**：  - 10 - 15 - 20
	TriggerAvailablePercent *ModifyAutoExpandPolicyReqTriggerAvailablePercent `json:"trigger_available_percent,omitempty"`

	// **参数解释**：  扩容步长百分比。  **取值范围**:   5 - 50
	StepPercent *int32 `json:"step_percent,omitempty"`
}

func (o ModifyAutoExpandPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAutoExpandPolicyReq struct{}"
	}

	return strings.Join([]string{"ModifyAutoExpandPolicyReq", string(data)}, " ")
}

type ModifyAutoExpandPolicyReqTriggerAvailablePercent struct {
	value int32
}

type ModifyAutoExpandPolicyReqTriggerAvailablePercentEnum struct {
	E_10 ModifyAutoExpandPolicyReqTriggerAvailablePercent
	E_15 ModifyAutoExpandPolicyReqTriggerAvailablePercent
	E_20 ModifyAutoExpandPolicyReqTriggerAvailablePercent
}

func GetModifyAutoExpandPolicyReqTriggerAvailablePercentEnum() ModifyAutoExpandPolicyReqTriggerAvailablePercentEnum {
	return ModifyAutoExpandPolicyReqTriggerAvailablePercentEnum{
		E_10: ModifyAutoExpandPolicyReqTriggerAvailablePercent{
			value: 10,
		}, E_15: ModifyAutoExpandPolicyReqTriggerAvailablePercent{
			value: 15,
		}, E_20: ModifyAutoExpandPolicyReqTriggerAvailablePercent{
			value: 20,
		},
	}
}

func (c ModifyAutoExpandPolicyReqTriggerAvailablePercent) Value() int32 {
	return c.value
}

func (c ModifyAutoExpandPolicyReqTriggerAvailablePercent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyAutoExpandPolicyReqTriggerAvailablePercent) UnmarshalJSON(b []byte) error {
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
