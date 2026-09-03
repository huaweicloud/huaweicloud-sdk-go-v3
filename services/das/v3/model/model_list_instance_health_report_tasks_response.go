package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceHealthReportTasksResponse Response Object
type ListInstanceHealthReportTasksResponse struct {

	// 诊断报告总数
	Total *int64 `json:"total,omitempty"`

	// 诊断报告列表
	HealthReportTaskList *[]HealthReportTaskInfo `json:"health_report_task_list,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ListInstanceHealthReportTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceHealthReportTasksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceHealthReportTasksResponse", string(data)}, " ")
}
