package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDesktopUsageMetricNewResponse Response Object
type ExportDesktopUsageMetricNewResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportDesktopUsageMetricNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesktopUsageMetricNewResponse struct{}"
	}

	return strings.Join([]string{"ExportDesktopUsageMetricNewResponse", string(data)}, " ")
}
