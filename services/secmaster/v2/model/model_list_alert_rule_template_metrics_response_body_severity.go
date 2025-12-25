package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleTemplateMetricsResponseBodySeverity 告警等级统计
type ListAlertRuleTemplateMetricsResponseBodySeverity struct {
	Metric *ListAlertRuleTemplateMetricsResponseBodySeverityMetric `json:"metric,omitempty"`

	// 分类方式
	Category *string `json:"category,omitempty"`
}

func (o ListAlertRuleTemplateMetricsResponseBodySeverity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplateMetricsResponseBodySeverity struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplateMetricsResponseBodySeverity", string(data)}, " ")
}
