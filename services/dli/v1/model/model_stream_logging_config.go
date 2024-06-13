package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StreamLoggingConfig 流日志配置。
type StreamLoggingConfig struct {

	// 是否开启作业的日志上传到用户的 OBS 功能。默认为 false。
	LogEnabled *bool `json:"log_enabled,omitempty"`

	// 用户授权保存作业日志的 OBS 桶名。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// 根目录日志级别配置，DEBUG\\TRACE\\WARNNING\\INFO\\ERROR
	RootLoggerLevel *string `json:"root_logger_level,omitempty"`

	// 输出日志的类名称对应的日志级别配置。
	LoggersLevelOfClass *[]StreamClassLoggerLevel `json:"loggers_level_of_class,omitempty"`
}

func (o StreamLoggingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamLoggingConfig struct{}"
	}

	return strings.Join([]string{"StreamLoggingConfig", string(data)}, " ")
}
