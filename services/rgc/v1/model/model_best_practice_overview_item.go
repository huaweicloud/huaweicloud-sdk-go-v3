package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BestPracticeOverviewItem 最佳实践检测结果中某个场景的概览。
type BestPracticeOverviewItem struct {

	// 分数
	Score float32 `json:"score,omitempty"`

	// 检测项个数
	DetectionCount *int32 `json:"detection_count,omitempty"`

	// 高风险项个数
	HighRiskCount *int32 `json:"high_risk_count,omitempty"`

	// 中风险项个数
	MediumRiskCount *int32 `json:"medium_risk_count,omitempty"`

	// 低风险项个数
	LowRiskCount *int32 `json:"low_risk_count,omitempty"`

	// 风险描述
	RiskItemDescription *[]string `json:"risk_item_description,omitempty"`
}

func (o BestPracticeOverviewItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BestPracticeOverviewItem struct{}"
	}

	return strings.Join([]string{"BestPracticeOverviewItem", string(data)}, " ")
}
