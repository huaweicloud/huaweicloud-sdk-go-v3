package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleTemplateMetricsResponseBodySeverityMetric 等级统计详情
type ListAlertRuleTemplateMetricsResponseBodySeverityMetric struct {

	// HIGH等级数量
	High *int32 `json:"HIGH,omitempty"`

	// TIPS等级数量
	Tips *int32 `json:"TIPS,omitempty"`

	// FATAL等级数量
	Fatal *int32 `json:"FATAL,omitempty"`

	// LOW等级数量
	Low *int32 `json:"LOW,omitempty"`

	// MEDIUM等级数量
	Medium *int32 `json:"MEDIUM,omitempty"`
}

func (o ListAlertRuleTemplateMetricsResponseBodySeverityMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplateMetricsResponseBodySeverityMetric struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplateMetricsResponseBodySeverityMetric", string(data)}, " ")
}
