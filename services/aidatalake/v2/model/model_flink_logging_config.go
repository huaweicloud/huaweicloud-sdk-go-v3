package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkLoggingConfig **参数解释**：Flink作业日志配置，可设置根目录日志级别和自定义日志类级别。 **约束限制**：不涉及。
type FlinkLoggingConfig struct {
	RootLoggerLevel *FlinkLoggerLevel `json:"root_logger_level,omitempty"`

	// **参数解释**：输出日志类的名称和对应的日志级别配置，用于为指定类设置特定的日志级别。 **约束限制**：配置项数量不能超过20个。
	Log4jLoggers *[]FlinkLog4jLoggers `json:"log4j_loggers,omitempty"`
}

func (o FlinkLoggingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkLoggingConfig struct{}"
	}

	return strings.Join([]string{"FlinkLoggingConfig", string(data)}, " ")
}
