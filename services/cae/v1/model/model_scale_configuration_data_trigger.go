package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScaleConfigurationDataTrigger 伸缩策略触发器。
type ScaleConfigurationDataTrigger struct {

	// 指标类型。
	Type *ScaleConfigurationDataTriggerType `json:"type,omitempty"`

	Metadata *ScalingTriggerMeta `json:"metadata,omitempty"`
}

func (o ScaleConfigurationDataTrigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleConfigurationDataTrigger struct{}"
	}

	return strings.Join([]string{"ScaleConfigurationDataTrigger", string(data)}, " ")
}

type ScaleConfigurationDataTriggerType struct {
	value string
}

type ScaleConfigurationDataTriggerTypeEnum struct {
	CPU    ScaleConfigurationDataTriggerType
	MEMORY ScaleConfigurationDataTriggerType
	CRON   ScaleConfigurationDataTriggerType
}

func GetScaleConfigurationDataTriggerTypeEnum() ScaleConfigurationDataTriggerTypeEnum {
	return ScaleConfigurationDataTriggerTypeEnum{
		CPU: ScaleConfigurationDataTriggerType{
			value: "cpu",
		},
		MEMORY: ScaleConfigurationDataTriggerType{
			value: "memory",
		},
		CRON: ScaleConfigurationDataTriggerType{
			value: "cron",
		},
	}
}

func (c ScaleConfigurationDataTriggerType) Value() string {
	return c.value
}

func (c ScaleConfigurationDataTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScaleConfigurationDataTriggerType) UnmarshalJSON(b []byte) error {
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
