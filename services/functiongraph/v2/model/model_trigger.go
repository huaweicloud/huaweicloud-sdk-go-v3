package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 触发器结构体
type Trigger struct {

	// 触发器名称
	TriggerName *string `json:"trigger_name,omitempty"`

	// 触发器类型
	TriggerType *TriggerTriggerType `json:"trigger_type,omitempty"`

	// 是否启用
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
	OBS       TriggerTriggerType
}

func GetTriggerTriggerTypeEnum() TriggerTriggerTypeEnum {
	return TriggerTriggerTypeEnum{
		FLOWTIMER: TriggerTriggerType{
			value: "FLOWTIMER",
		},
		OBS: TriggerTriggerType{
			value: "OBS",
		},
	}
}

func (c TriggerTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerTriggerType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
