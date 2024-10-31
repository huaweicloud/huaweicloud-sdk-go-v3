package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportRiskSuggestion struct {

	// 建议优化措施编码。
	SuggestionCode string `json:"suggestion_code"`

	// 建议优化措施。
	SuggestionContent string `json:"suggestion_content"`
}

func (o HealthReportRiskSuggestion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportRiskSuggestion struct{}"
	}

	return strings.Join([]string{"HealthReportRiskSuggestion", string(data)}, " ")
}
