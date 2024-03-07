package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricNotifyRuleRequest Request Object
type ListMetricNotifyRuleRequest struct {

	// 指标名称(精确匹配)
	MetricName *string `json:"metric_name,omitempty"`

	// 通知规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 查询的偏移量,默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100],默认值10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMetricNotifyRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricNotifyRuleRequest struct{}"
	}

	return strings.Join([]string{"ListMetricNotifyRuleRequest", string(data)}, " ")
}
