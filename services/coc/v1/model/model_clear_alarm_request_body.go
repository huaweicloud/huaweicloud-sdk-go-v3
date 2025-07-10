package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClearAlarmRequestBody struct {

	// 清除备注
	Remarks *string `json:"remarks,omitempty"`

	// 告警id拼接字符串，以”，“分隔
	AlarmIds string `json:"alarm_ids"`

	// 业务是否中断
	IsServiceInterrupt *bool `json:"is_service_interrupt,omitempty"`

	// 故障发生时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 故障恢复时间
	FaultRecoveryTime *int64 `json:"fault_recovery_time,omitempty"`
}

func (o ClearAlarmRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearAlarmRequestBody struct{}"
	}

	return strings.Join([]string{"ClearAlarmRequestBody", string(data)}, " ")
}
