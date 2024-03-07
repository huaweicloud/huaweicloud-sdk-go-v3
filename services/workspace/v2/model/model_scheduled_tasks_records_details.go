package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduledTasksRecordsDetails 定时任务执行记录详情。
type ScheduledTasksRecordsDetails struct {

	// 任务执行记录详情id
	Id *string `json:"id,omitempty"`

	// 任务执行记录id
	RecordId *string `json:"record_id,omitempty"`

	// 桌面id
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 执行状态。
	ExecStatus *string `json:"exec_status,omitempty"`

	// 失败或者跳过原因的错误码。
	ResultCode *string `json:"result_code,omitempty"`

	// 失败或者跳过原因。
	FailReason *string `json:"fail_reason,omitempty"`

	// 执行开始时间，格式为yyyy-MM-dd HH:mm:ss。
	StartTime *string `json:"start_time,omitempty"`

	// 执行结束时间，格式为yyyy-MM-dd HH:mm:ss。
	EndTime *string `json:"end_time,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o ScheduledTasksRecordsDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTasksRecordsDetails struct{}"
	}

	return strings.Join([]string{"ScheduledTasksRecordsDetails", string(data)}, " ")
}
