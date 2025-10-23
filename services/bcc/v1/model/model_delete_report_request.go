package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReportRequest Request Object
type DeleteReportRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 报告ID
	ReportId string `json:"report_id"`
}

func (o DeleteReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReportRequest struct{}"
	}

	return strings.Join([]string{"DeleteReportRequest", string(data)}, " ")
}
