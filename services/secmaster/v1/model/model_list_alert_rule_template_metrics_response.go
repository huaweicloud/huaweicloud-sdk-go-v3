package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleTemplateMetricsResponse Response Object
type ListAlertRuleTemplateMetricsResponse struct {

	// 响应结果
	Body           map[string]AlertRuleTemplateMetric `json:"body,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListAlertRuleTemplateMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplateMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplateMetricsResponse", string(data)}, " ")
}
