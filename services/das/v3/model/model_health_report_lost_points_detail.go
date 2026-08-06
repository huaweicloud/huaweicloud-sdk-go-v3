package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportLostPointsDetail struct {

	// 扣分项。
	Metric *string `json:"metric,omitempty"`

	// 所扣分数。
	LostPoints *float64 `json:"lost_points,omitempty"`

	// 事件等级。
	RiskLevel *string `json:"risk_level,omitempty"`
}

func (o HealthReportLostPointsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportLostPointsDetail struct{}"
	}

	return strings.Join([]string{"HealthReportLostPointsDetail", string(data)}, " ")
}
