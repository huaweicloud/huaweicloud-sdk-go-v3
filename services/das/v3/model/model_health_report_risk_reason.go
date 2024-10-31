package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportRiskReason struct {

	// 可能原因编码。
	ReasonCode string `json:"reason_code"`

	// 可能原因内容。
	ReasonContent string `json:"reason_content"`

	// 建议优化措施列表。
	Suggestions []HealthReportRiskSuggestion `json:"suggestions"`
}

func (o HealthReportRiskReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportRiskReason struct{}"
	}

	return strings.Join([]string{"HealthReportRiskReason", string(data)}, " ")
}
