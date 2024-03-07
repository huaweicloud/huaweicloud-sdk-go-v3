package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricNotifyRuleResponse Response Object
type ListMetricNotifyRuleResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 规则列表
	Items          *[]MetricNotifyRule `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListMetricNotifyRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricNotifyRuleResponse struct{}"
	}

	return strings.Join([]string{"ListMetricNotifyRuleResponse", string(data)}, " ")
}
