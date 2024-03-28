package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkJobStatusReportRequest Request Object
type UpdateFlinkJobStatusReportRequest struct {
	Body *ChangeFlinkJobStatusReportRequestBody `json:"body,omitempty"`
}

func (o UpdateFlinkJobStatusReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkJobStatusReportRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlinkJobStatusReportRequest", string(data)}, " ")
}
