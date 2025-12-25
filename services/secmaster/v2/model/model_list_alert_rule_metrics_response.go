package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleMetricsResponse Response Object
type ListAlertRuleMetricsResponse struct {
	CuUsage *CuUsage `json:"cu_usage,omitempty"`

	AlertSeverities *AlertSeverities `json:"alert_severities,omitempty"`

	MetricsStatus  *MetricsStatus `json:"metrics_status,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListAlertRuleMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListAlertRuleMetricsResponse", string(data)}, " ")
}
