package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecyclePolicy 回收站策略。
type RecyclePolicy struct {

	// 保留天数，1-7天。
	RetentionPeriodInDays string `json:"retention_period_in_days"`
}

func (o RecyclePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecyclePolicy struct{}"
	}

	return strings.Join([]string{"RecyclePolicy", string(data)}, " ")
}
