package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StatusSchema **参数解释** 资源分组健康状态 **约束限制** 不涉及 **取值范围** - health: 表示无告警 - unhealth: 表示告警中 - no_alarm_rule: 表示未设置告警规则 **默认取值** 不涉及
type StatusSchema struct {
	value string
}

type StatusSchemaEnum struct {
	HEALTH        StatusSchema
	UNHEALTH      StatusSchema
	NO_ALARM_RULE StatusSchema
}

func GetStatusSchemaEnum() StatusSchemaEnum {
	return StatusSchemaEnum{
		HEALTH: StatusSchema{
			value: "health",
		},
		UNHEALTH: StatusSchema{
			value: "unhealth",
		},
		NO_ALARM_RULE: StatusSchema{
			value: "no_alarm_rule",
		},
	}
}

func (c StatusSchema) Value() string {
	return c.value
}

func (c StatusSchema) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatusSchema) UnmarshalJSON(b []byte) error {
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
