package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StatusSchemaResp **参数解释** 资源分组健康状态 **取值范围** - health: 表示无告警 - unhealth: 表示告警中 - no_alarm_rule: 表示未设置告警规则
type StatusSchemaResp struct {
	value string
}

type StatusSchemaRespEnum struct {
	HEALTH        StatusSchemaResp
	UNHEALTH      StatusSchemaResp
	NO_ALARM_RULE StatusSchemaResp
}

func GetStatusSchemaRespEnum() StatusSchemaRespEnum {
	return StatusSchemaRespEnum{
		HEALTH: StatusSchemaResp{
			value: "health",
		},
		UNHEALTH: StatusSchemaResp{
			value: "unhealth",
		},
		NO_ALARM_RULE: StatusSchemaResp{
			value: "no_alarm_rule",
		},
	}
}

func (c StatusSchemaResp) Value() string {
	return c.value
}

func (c StatusSchemaResp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatusSchemaResp) UnmarshalJSON(b []byte) error {
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
