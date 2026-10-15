package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkLog4jLoggersResponseBody **参数解释**：类日志级别配置，包含指定类的日志级别配置。
type ShowFlinkLog4jLoggersResponseBody struct {

	// **参数解释**：日志类名称。 **取值范围**：1~256个字符。
	LoggerName string `json:"logger_name"`

	LoggerLevel *FlinkLoggerLevel `json:"logger_level"`
}

func (o ShowFlinkLog4jLoggersResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkLog4jLoggersResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFlinkLog4jLoggersResponseBody", string(data)}, " ")
}
