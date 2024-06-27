package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertStatisticsDto struct {

	// 阻塞告警次数
	BlockAlertCount *int32 `json:"block_alert_count,omitempty"`

	// 异常告警次数
	ExceptionAlertCount *int32 `json:"exception_alert_count,omitempty"`

	// 失败告警次数
	FailAlertCount *int32 `json:"fail_alert_count,omitempty"`

	// 服务id
	ServiceId *string `json:"service_id,omitempty"`

	// 统计时间
	StatisticsTime *int64 `json:"statistics_time,omitempty"`

	// 超时告警次数
	TimeoutAlertCount *int32 `json:"timeout_alert_count,omitempty"`
}

func (o AlertStatisticsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertStatisticsDto struct{}"
	}

	return strings.Join([]string{"AlertStatisticsDto", string(data)}, " ")
}
