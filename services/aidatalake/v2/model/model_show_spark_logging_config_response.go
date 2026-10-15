package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkLoggingConfigResponse **参数解释**：日志配置信息，包含Driver和Executor的日志级别配置。
type ShowSparkLoggingConfigResponse struct {
	DriverRootLoggerLevel *LogLevel `json:"driver_root_logger_level,omitempty"`

	// **参数解释**：Driver类日志级别配置列表，用于为指定类设置特定的日志级别。
	DriverLoggersLevelOfClass *[]ShowSparkClassLoggerLevelResponse `json:"driver_loggers_level_of_class,omitempty"`

	ExecutorRootLoggerLevel *LogLevel `json:"executor_root_logger_level,omitempty"`

	// **参数解释**：Executor类日志级别配置列表，用于为指定类设置特定的日志级别。
	ExecutorLoggersLevelOfClass *[]ShowSparkClassLoggerLevelResponse `json:"executor_loggers_level_of_class,omitempty"`
}

func (o ShowSparkLoggingConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkLoggingConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkLoggingConfigResponse", string(data)}, " ")
}
