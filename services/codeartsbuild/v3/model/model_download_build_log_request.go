package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadBuildLogRequest Request Object
type DownloadBuildLogRequest struct {

	// 记录ID,36位数字、小写字母、'-'组组合。
	RecordId string `json:"record_id"`

	// 日志等级 值为INFO | DEBUG。
	LogLevel *string `json:"log_level,omitempty"`
}

func (o DownloadBuildLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBuildLogRequest struct{}"
	}

	return strings.Join([]string{"DownloadBuildLogRequest", string(data)}, " ")
}
