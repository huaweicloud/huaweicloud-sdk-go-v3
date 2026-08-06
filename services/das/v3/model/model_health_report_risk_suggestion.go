package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthReportRiskSuggestion 建议优化操作
type HealthReportRiskSuggestion struct {

	// 建议优化措施编码
	SuggestionCode *string `json:"suggestion_code,omitempty"`

	// 建议优化措施
	SuggestionContent *string `json:"suggestion_content,omitempty"`
}

func (o HealthReportRiskSuggestion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportRiskSuggestion struct{}"
	}

	return strings.Join([]string{"HealthReportRiskSuggestion", string(data)}, " ")
}
