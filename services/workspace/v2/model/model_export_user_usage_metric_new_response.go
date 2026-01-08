package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserUsageMetricNewResponse Response Object
type ExportUserUsageMetricNewResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportUserUsageMetricNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserUsageMetricNewResponse struct{}"
	}

	return strings.Join([]string{"ExportUserUsageMetricNewResponse", string(data)}, " ")
}
