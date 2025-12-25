package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleTemplateMetricsResponseBodyStatusMetric 状态详情
type ListAlertRuleTemplateMetricsResponseBodyStatusMetric struct {

	// 活跃状态数量
	Activity *int32 `json:"activity,omitempty"`

	// 启用状态数量
	Enabled *int32 `json:"ENABLED,omitempty"`

	// 禁用状态数量
	Disabled *int32 `json:"DISABLED,omitempty"`
}

func (o ListAlertRuleTemplateMetricsResponseBodyStatusMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplateMetricsResponseBodyStatusMetric struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplateMetricsResponseBodyStatusMetric", string(data)}, " ")
}
