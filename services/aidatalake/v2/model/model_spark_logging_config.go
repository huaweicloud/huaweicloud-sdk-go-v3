package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkLoggingConfig **参数解释**：Spark作业日志配置，用于配置Driver和Executor的日志级别。 **约束限制**：不涉及。
type SparkLoggingConfig struct {
	DriverRootLoggerLevel *LogLevel `json:"driver_root_logger_level,omitempty"`

	// **参数解释**：Driver类级别日志配置列表，用于配置指定类的日志级别。数组中的每个元素为SparkClassLoggerLevel对象，包含类名和日志级别。 **约束限制**：最多配置20个类的日志级别。
	DriverLoggersLevelOfClass *[]SparkClassLoggerLevel `json:"driver_loggers_level_of_class,omitempty"`

	ExecutorRootLoggerLevel *LogLevel `json:"executor_root_logger_level,omitempty"`

	// **参数解释**：Executor类级别日志配置列表，用于配置指定类的日志级别。数组中的每个元素为SparkClassLoggerLevel对象，包含类名和日志级别。 **约束限制**：最多配置20个类的日志级别。
	ExecutorLoggersLevelOfClass *[]SparkClassLoggerLevel `json:"executor_loggers_level_of_class,omitempty"`
}

func (o SparkLoggingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkLoggingConfig struct{}"
	}

	return strings.Join([]string{"SparkLoggingConfig", string(data)}, " ")
}
