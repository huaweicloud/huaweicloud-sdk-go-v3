package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchInspectionLostPointsDetail struct {

	// 风险等级
	RiskLevel string `json:"risk_level"`

	// 检查项
	Metric string `json:"metric"`

	// 指标值
	MetricValue float64 `json:"metric_value"`

	// 扣分值
	DeductedPoints float64 `json:"deducted_points"`

	// 扣分条件
	DeductedCondition string `json:"deducted_condition"`

	// 扣分规则
	DeductedFormula string `json:"deducted_formula"`

	// 优化建议
	Suggestions string `json:"suggestions"`
}

func (o BatchInspectionLostPointsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInspectionLostPointsDetail struct{}"
	}

	return strings.Join([]string{"BatchInspectionLostPointsDetail", string(data)}, " ")
}
