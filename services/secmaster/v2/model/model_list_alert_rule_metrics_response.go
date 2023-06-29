package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleMetricsResponse Response Object
type ListAlertRuleMetricsResponse struct {

	// ListAlertRuleMetricsResponseBody
	Body map[string]AlertRuleMetric `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAlertRuleMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListAlertRuleMetricsResponse", string(data)}, " ")
}
