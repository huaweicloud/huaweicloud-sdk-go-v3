package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObejectLevelCompareOverviewResponse Response Object
type ListObejectLevelCompareOverviewResponse struct {

	// 对比任务创建时间，UTC时间，例如：2024-04-09T07:00:57Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 对比任务开始时间，UTC时间，例如：2024-04-09T07:00:57Z。
	StartTime *string `json:"start_time,omitempty"`

	// 对比任务状态。 取值： - RUNNING：运行中。 - WAITING_FOR_RUNNING：等待启动中。 - SUCCESSFUL：完成。 - FAILED：失败。 - CANCELLED：已取消。 - TIMEOUT_INTERRUPT：超时中断。 - FULL_DOING：全量校验中。 - INCRE_DOING：增量校验中。
	Status *string `json:"status,omitempty"`

	// 生成对比结果报告文件的状态：  - INIT：初始状态。  - EXPORTING：对比结果导出中。  - EXPORT_COMPLETE：对比结果导出完成。  - EXPORT_COMMON_FAILED：对比结果导出失败。
	ExportStatus *string `json:"export_status,omitempty"`

	// 对比结果报告文件有效期剩余时间，单位秒，报告未生成时返回-1。
	ReportRemainSeconds *int64 `json:"report_remain_seconds,omitempty"`

	// 对比任务ID。
	CompareJobId *string `json:"compare_job_id,omitempty"`

	// 失败原因。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 对比结果。
	CompareResult  *[]ObjectsCompareOverviewInfo `json:"compare_result,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListObejectLevelCompareOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObejectLevelCompareOverviewResponse struct{}"
	}

	return strings.Join([]string{"ListObejectLevelCompareOverviewResponse", string(data)}, " ")
}
