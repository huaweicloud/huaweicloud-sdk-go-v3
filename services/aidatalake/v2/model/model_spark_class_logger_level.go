package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkClassLoggerLevel **参数解释**：类级别日志配置，用于配置指定类的日志级别。 **约束限制**：不涉及。
type SparkClassLoggerLevel struct {

	// **参数解释**：日志类名称，用于指定需要配置日志级别的类名。 **约束限制**：不涉及。 **取值范围**：长度为1~256个字符。 **默认取值**：不涉及。
	LoggerName *string `json:"logger_name,omitempty"`

	LoggerLevel *LogLevel `json:"logger_level,omitempty"`
}

func (o SparkClassLoggerLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkClassLoggerLevel struct{}"
	}

	return strings.Join([]string{"SparkClassLoggerLevel", string(data)}, " ")
}
