package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkSavepointTriggerType **参数解释**：Savepoint触发类型枚举。 **取值范围**：固定值USER_REQUEST，表示由用户手动生成。
type FlinkSavepointTriggerType struct {
	value string
}

type FlinkSavepointTriggerTypeEnum struct {
	USER_REQUEST FlinkSavepointTriggerType
}

func GetFlinkSavepointTriggerTypeEnum() FlinkSavepointTriggerTypeEnum {
	return FlinkSavepointTriggerTypeEnum{
		USER_REQUEST: FlinkSavepointTriggerType{
			value: "USER_REQUEST",
		},
	}
}

func (c FlinkSavepointTriggerType) Value() string {
	return c.value
}

func (c FlinkSavepointTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkSavepointTriggerType) UnmarshalJSON(b []byte) error {
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
