package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkLoggingConfigResponseBody **参数解释**：Flink作业日志配置，包含根目录日志级别和自定义日志类级别。
type ShowFlinkLoggingConfigResponseBody struct {
	RootLoggerLevel *FlinkLoggerLevel `json:"root_logger_level,omitempty"`

	// **参数解释**：自定义日志类配置列表。
	Log4jLoggers *[]ShowFlinkLog4jLoggersResponseBody `json:"log4j_loggers,omitempty"`
}

func (o ShowFlinkLoggingConfigResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkLoggingConfigResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFlinkLoggingConfigResponseBody", string(data)}, " ")
}
