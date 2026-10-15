package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkLog4jLoggers **参数解释**：类日志级别配置，用于为指定类设置特定的日志级别。 **约束限制**：不涉及。
type FlinkLog4jLoggers struct {

	// **参数解释**：输出日志类的名称，用于指定需要配置日志级别的类。 **约束限制**：不涉及。 **取值范围**：1~256个字符，字符组成不限制。 **默认取值**：不涉及。
	LoggerName string `json:"logger_name"`

	LoggerLevel *FlinkLoggerLevel `json:"logger_level"`
}

func (o FlinkLog4jLoggers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkLog4jLoggers struct{}"
	}

	return strings.Join([]string{"FlinkLog4jLoggers", string(data)}, " ")
}
