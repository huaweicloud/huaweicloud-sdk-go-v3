package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AbuseThreshold 算力滥用触发阈值
type AbuseThreshold struct {

	// 统计时间窗口（秒），最大86400（1天）
	TimeWindowSeconds *int32 `json:"time_window_seconds,omitempty"`

	// 阈值条件列表，当前仅支持1个条件
	Conditions *[]AbuseCondition `json:"conditions,omitempty"`
}

func (o AbuseThreshold) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AbuseThreshold struct{}"
	}

	return strings.Join([]string{"AbuseThreshold", string(data)}, " ")
}
