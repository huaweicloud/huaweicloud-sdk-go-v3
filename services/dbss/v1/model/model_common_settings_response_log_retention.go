package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommonSettingsResponseLogRetention 日志保留设置
type CommonSettingsResponseLogRetention struct {

	// 设定的审计日志保存时间
	RetentionDays *int32 `json:"retention_days,omitempty"`

	// 审计日志保存时间设置最小时间
	RangeDaysMin *int32 `json:"range_days_min,omitempty"`

	// 审计日志保存时间设置最大时间
	RangeDaysMax *int32 `json:"range_days_max,omitempty"`
}

func (o CommonSettingsResponseLogRetention) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonSettingsResponseLogRetention struct{}"
	}

	return strings.Join([]string{"CommonSettingsResponseLogRetention", string(data)}, " ")
}
