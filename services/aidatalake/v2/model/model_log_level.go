package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LogLevel **参数解释**：日志级别配置，用于控制日志输出的详细程度。 **约束限制**：不涉及。 **取值范围**： - DEBUG：调试级别。 - TRACE：跟踪级别。 - WARN：警告级别。 - INFO：信息级别。 - ERROR：错误级别。 **默认取值**：INFO。
type LogLevel struct {
	value string
}

type LogLevelEnum struct {
	DEBUG LogLevel
	TRACE LogLevel
	WARN  LogLevel
	INFO  LogLevel
	ERROR LogLevel
}

func GetLogLevelEnum() LogLevelEnum {
	return LogLevelEnum{
		DEBUG: LogLevel{
			value: "DEBUG",
		},
		TRACE: LogLevel{
			value: "TRACE",
		},
		WARN: LogLevel{
			value: "WARN",
		},
		INFO: LogLevel{
			value: "INFO",
		},
		ERROR: LogLevel{
			value: "ERROR",
		},
	}
}

func (c LogLevel) Value() string {
	return c.value
}

func (c LogLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LogLevel) UnmarshalJSON(b []byte) error {
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
