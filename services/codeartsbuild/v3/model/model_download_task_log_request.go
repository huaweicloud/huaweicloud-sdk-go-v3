package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadTaskLogRequest Request Object
type DownloadTaskLogRequest struct {

	// 记录ID,36位数字、小写字母、'-'组组合。
	RecordId string `json:"record_id"`

	// 步骤名称
	TaskName string `json:"task_name"`

	// 日志等级 值为INFO | DEBUG。
	LogLevel *string `json:"log_level,omitempty"`
}

func (o DownloadTaskLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadTaskLogRequest struct{}"
	}

	return strings.Join([]string{"DownloadTaskLogRequest", string(data)}, " ")
}
