package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StreamClassLoggerLevel 输出日志的类名称对应的日志级别配置。
type StreamClassLoggerLevel struct {

	// 输出日志的类的名称。
	LoggerName *string `json:"logger_name,omitempty"`

	// 输出日志的级别，DEBUG\\TRACE\\WARNNING\\INFO\\ERROR。
	LoggerLevel *string `json:"logger_level,omitempty"`
}

func (o StreamClassLoggerLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamClassLoggerLevel struct{}"
	}

	return strings.Join([]string{"StreamClassLoggerLevel", string(data)}, " ")
}
