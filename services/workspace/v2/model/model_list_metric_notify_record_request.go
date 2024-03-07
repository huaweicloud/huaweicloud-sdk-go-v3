package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricNotifyRecordRequest Request Object
type ListMetricNotifyRecordRequest struct {

	// 指标名称(精确匹配) (metric_name和rule_id不允许同时为空)
	MetricName *string `json:"metric_name,omitempty"`

	// 通知规则ID (metric_name和rule_id不允许同时为空)
	RuleId *string `json:"rule_id,omitempty"`

	// 查询的偏移量,默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100],默认值10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMetricNotifyRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricNotifyRecordRequest struct{}"
	}

	return strings.Join([]string{"ListMetricNotifyRecordRequest", string(data)}, " ")
}
