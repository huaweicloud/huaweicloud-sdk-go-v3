package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Trigger 触发器结构体
type Trigger struct {

	// 触发器名称
	TriggerName string `json:"trigger_name"`

	// 触发器类型 FLOWTIMER：定时触发器 SMN：SMN触发器 APIG：APIG触发器(共享版) APIG_DE：APIG触发器(专享版) OBS：OBS触发器
	TriggerType TriggerTriggerType `json:"trigger_type"`

	// 是否启用触发器
	Enabled *bool `json:"enabled,omitempty"`

	TriggerConfig *ObsTriggerConfig `json:"trigger_config,omitempty"`
}

func (o Trigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Trigger struct{}"
	}

	return strings.Join([]string{"Trigger", string(data)}, " ")
}

type TriggerTriggerType struct {
	value string
}

type TriggerTriggerTypeEnum struct {
	FLOWTIMER TriggerTriggerType
	SMN       TriggerTriggerType
	APIG      TriggerTriggerType
	APIG_DE   TriggerTriggerType
	OBS       TriggerTriggerType
}

func GetTriggerTriggerTypeEnum() TriggerTriggerTypeEnum {
	return TriggerTriggerTypeEnum{
		FLOWTIMER: TriggerTriggerType{
			value: "FLOWTIMER",
		},
		SMN: TriggerTriggerType{
			value: "SMN",
		},
		APIG: TriggerTriggerType{
			value: "APIG",
		},
		APIG_DE: TriggerTriggerType{
			value: "APIG_DE",
		},
		OBS: TriggerTriggerType{
			value: "OBS",
		},
	}
}

func (c TriggerTriggerType) Value() string {
	return c.value
}

func (c TriggerTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerTriggerType) UnmarshalJSON(b []byte) error {
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
