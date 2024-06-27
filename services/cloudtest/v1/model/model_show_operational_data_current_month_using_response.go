package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOperationalDataCurrentMonthUsingResponse Response Object
type ShowOperationalDataCurrentMonthUsingResponse struct {

	// 告警成功比率，恒为1
	AlertSuccessRate *int32 `json:"alert_success_rate,omitempty"`

	// 正在运行的任务个数
	RunningTasks *int32 `json:"running_tasks,omitempty"`

	// 总告警数
	TotalAlerts *int32 `json:"total_alerts,omitempty"`

	// 总运行天数
	TotalDays *int32 `json:"total_days,omitempty"`

	// 总运行个数
	TotalRuns *int64 `json:"total_runs,omitempty"`

	// 总任务个数
	TotalTasks *int32 `json:"total_tasks,omitempty"`

	// 工作项个数
	WorkItemCount  *int32 `json:"work_item_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowOperationalDataCurrentMonthUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOperationalDataCurrentMonthUsingResponse struct{}"
	}

	return strings.Join([]string{"ShowOperationalDataCurrentMonthUsingResponse", string(data)}, " ")
}
