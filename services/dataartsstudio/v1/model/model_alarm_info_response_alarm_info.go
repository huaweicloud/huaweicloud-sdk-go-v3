package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmInfoResponseAlarmInfo struct {

	// 告警通知时间
	AlarmTime *int64 `json:"alarm_time,omitempty"`

	// 作业名称
	JobName *string `json:"job_name,omitempty"`

	// 作业实例调度方式，取值范围：0 正常调度，2手工调度，5补数据，6子作业调度，7单次调度
	ScheduleType *int32 `json:"schedule_type,omitempty"`

	// 发送信息
	SendMsg *string `json:"send_msg,omitempty"`

	// 计划时间
	PlanTime *int64 `json:"plan_time,omitempty"`

	// 告警通知类型，取值范围：0 运行成功，1 运行异常/失败， 12 未完成，13 运行取消
	RemindType *int32 `json:"remind_type,omitempty"`

	// 发送状态，取值范围：0 发送成功，1：发送失败
	SendStatus *int32 `json:"send_status,omitempty"`

	// 作业ID
	JobId *int64 `json:"job_id,omitempty"`
}

func (o AlarmInfoResponseAlarmInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmInfoResponseAlarmInfo struct{}"
	}

	return strings.Join([]string{"AlarmInfoResponseAlarmInfo", string(data)}, " ")
}
