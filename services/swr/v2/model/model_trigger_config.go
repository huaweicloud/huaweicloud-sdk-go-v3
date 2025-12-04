package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TriggerConfig struct {

	// 触发类型，老化规则只支持manual(手动)、scheduled(定时+手动)；同步策略支持manual(手动)、scheduled(定时+手动)、event_based(事件触发+手动);镜像签名支持manual(手动)、event_based(事件触发+手动)
	Type TriggerConfigType `json:"type"`

	TriggerSettings *TriggerSetting `json:"trigger_settings,omitempty"`
}

func (o TriggerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerConfig struct{}"
	}

	return strings.Join([]string{"TriggerConfig", string(data)}, " ")
}

type TriggerConfigType struct {
	value string
}

type TriggerConfigTypeEnum struct {
	MANUAL      TriggerConfigType
	SCHEDULED   TriggerConfigType
	EVENT_BASED TriggerConfigType
}

func GetTriggerConfigTypeEnum() TriggerConfigTypeEnum {
	return TriggerConfigTypeEnum{
		MANUAL: TriggerConfigType{
			value: "manual",
		},
		SCHEDULED: TriggerConfigType{
			value: "scheduled",
		},
		EVENT_BASED: TriggerConfigType{
			value: "event_based",
		},
	}
}

func (c TriggerConfigType) Value() string {
	return c.value
}

func (c TriggerConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerConfigType) UnmarshalJSON(b []byte) error {
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
