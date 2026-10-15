package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkClassLoggerLevelResponse **参数解释**：类日志级别配置，用于为指定类设置特定的日志级别。
type ShowSparkClassLoggerLevelResponse struct {

	// **参数解释**：输出日志类的名称，用于指定需要配置日志级别的类。 **取值范围**：完整的类路径格式，例如：com.aaa.bbb.ccc。长度为1~512个字符。
	LoggerName *string `json:"logger_name,omitempty"`

	LoggerLevel *LogLevel `json:"logger_level,omitempty"`
}

func (o ShowSparkClassLoggerLevelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkClassLoggerLevelResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkClassLoggerLevelResponse", string(data)}, " ")
}
