package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleTaskExecuteDetail 定时任务执行记录子任务详情。
type ScheduleTaskExecuteDetail struct {

	// 定时任务执行记录主键id。
	Id *string `json:"id,omitempty"`

	// 主任务记录id。
	ExecuteId *string `json:"execute_id,omitempty"`

	// 操作的server_id。
	ServerId *string `json:"server_id,omitempty"`

	// 操作的机器名称。
	ServerName *string `json:"server_name,omitempty"`

	Status *ScheduleTaskStatus `json:"status,omitempty"`

	TaskType *ScheduleTaskTypeEnum `json:"task_type,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 子任务开始时间。
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 子任务结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 任务执行失败时的错误码。
	ResultCode *string `json:"result_code,omitempty"`

	// 任务失败原因。
	ResultMessage *string `json:"result_message,omitempty"`
}

func (o ScheduleTaskExecuteDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleTaskExecuteDetail struct{}"
	}

	return strings.Join([]string{"ScheduleTaskExecuteDetail", string(data)}, " ")
}
