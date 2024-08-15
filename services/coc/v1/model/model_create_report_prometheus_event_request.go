package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportPrometheusEventRequest Request Object
type CreateReportPrometheusEventRequest struct {

	// 集成ID
	IntegrationKey string `json:"integration_key"`

	Body map[string]interface{} `json:"body,omitempty"`
}

func (o CreateReportPrometheusEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportPrometheusEventRequest struct{}"
	}

	return strings.Join([]string{"CreateReportPrometheusEventRequest", string(data)}, " ")
}
