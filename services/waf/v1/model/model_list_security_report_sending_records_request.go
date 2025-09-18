package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityReportSendingRecordsRequest Request Object
type ListSecurityReportSendingRecordsRequest struct {

	// 报告名称
	ReportName *string `json:"report_name,omitempty"`

	// 报告类别
	ReportCategory *string `json:"report_category,omitempty"`

	// 限制条目数量
	Limit *int64 `json:"limit,omitempty"`

	// 偏移量
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListSecurityReportSendingRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityReportSendingRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityReportSendingRecordsRequest", string(data)}, " ")
}
