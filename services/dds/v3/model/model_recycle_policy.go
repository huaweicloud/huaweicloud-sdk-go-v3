package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecyclePolicy struct {

	// 打开回收策略，不可关闭 - true 打开回收策略
	Enabled bool `json:"enabled"`

	// 策略保持时长（1-7天），天数为正整数，不填默认保留7天
	RetentionPeriodInDays *int32 `json:"retention_period_in_days,omitempty"`
}

func (o RecyclePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecyclePolicy struct{}"
	}

	return strings.Join([]string{"RecyclePolicy", string(data)}, " ")
}
