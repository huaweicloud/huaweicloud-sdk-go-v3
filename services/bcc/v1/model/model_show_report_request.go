package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportRequest Request Object
type ShowReportRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 报告ID
	ReportId string `json:"report_id"`
}

func (o ShowReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportRequest struct{}"
	}

	return strings.Join([]string{"ShowReportRequest", string(data)}, " ")
}
