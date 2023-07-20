package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeFlinkJobStatusReportRequest Request Object
type ChangeFlinkJobStatusReportRequest struct {
	Body *IefFlinkJobStatusReportReq `json:"body,omitempty"`
}

func (o ChangeFlinkJobStatusReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFlinkJobStatusReportRequest struct{}"
	}

	return strings.Join([]string{"ChangeFlinkJobStatusReportRequest", string(data)}, " ")
}
