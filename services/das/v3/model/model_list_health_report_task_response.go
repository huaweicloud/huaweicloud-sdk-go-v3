package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHealthReportTaskResponse Response Object
type ListHealthReportTaskResponse struct {

	// 诊断报告总数
	Total *int64 `json:"total,omitempty"`

	// 诊断报告列表
	HealthReportTaskList *[]HealthReportTask `json:"health_report_task_list,omitempty"`
	HttpStatusCode       int                 `json:"-"`
}

func (o ListHealthReportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthReportTaskResponse struct{}"
	}

	return strings.Join([]string{"ListHealthReportTaskResponse", string(data)}, " ")
}
