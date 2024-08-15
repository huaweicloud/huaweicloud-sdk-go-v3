package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportCustomEventRequest Request Object
type CreateReportCustomEventRequest struct {

	// 集成ID
	IntegrationKey string `json:"integration_key"`

	Body *ReportCustomEventRequestBody `json:"body,omitempty"`
}

func (o CreateReportCustomEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportCustomEventRequest struct{}"
	}

	return strings.Join([]string{"CreateReportCustomEventRequest", string(data)}, " ")
}
