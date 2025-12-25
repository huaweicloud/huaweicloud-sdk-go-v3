package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleTemplateMetricsResponseBodyStatus 告警规则状态统计
type ListAlertRuleTemplateMetricsResponseBodyStatus struct {
	Metric *ListAlertRuleTemplateMetricsResponseBodyStatusMetric `json:"metric,omitempty"`

	// 分类方式
	Category *string `json:"category,omitempty"`
}

func (o ListAlertRuleTemplateMetricsResponseBodyStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplateMetricsResponseBodyStatus struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplateMetricsResponseBodyStatus", string(data)}, " ")
}
