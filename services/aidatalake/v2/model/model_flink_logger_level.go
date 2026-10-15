package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkLoggerLevel **参数解释**：Flink作业可控制的日志等级枚举。 **约束限制**：不涉及。 **取值范围**：   - TRACE：跟踪。   - DEBUG：调试。   - INFO：信息。   - WARN：警告。   - ERROR：错误。 **默认取值**：INFO。
type FlinkLoggerLevel struct {
	value string
}

type FlinkLoggerLevelEnum struct {
	TRACE FlinkLoggerLevel
	DEBUG FlinkLoggerLevel
	INFO  FlinkLoggerLevel
	WARN  FlinkLoggerLevel
	ERROR FlinkLoggerLevel
}

func GetFlinkLoggerLevelEnum() FlinkLoggerLevelEnum {
	return FlinkLoggerLevelEnum{
		TRACE: FlinkLoggerLevel{
			value: "TRACE",
		},
		DEBUG: FlinkLoggerLevel{
			value: "DEBUG",
		},
		INFO: FlinkLoggerLevel{
			value: "INFO",
		},
		WARN: FlinkLoggerLevel{
			value: "WARN",
		},
		ERROR: FlinkLoggerLevel{
			value: "ERROR",
		},
	}
}

func (c FlinkLoggerLevel) Value() string {
	return c.value
}

func (c FlinkLoggerLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkLoggerLevel) UnmarshalJSON(b []byte) error {
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
