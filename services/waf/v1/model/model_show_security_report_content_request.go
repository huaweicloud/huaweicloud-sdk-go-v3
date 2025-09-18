package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityReportContentRequest Request Object
type ShowSecurityReportContentRequest struct {

	// 报告ID
	ReportId string `json:"report_id"`

	// 订阅ID
	SubscriptionId string `json:"subscription_id"`
}

func (o ShowSecurityReportContentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityReportContentRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityReportContentRequest", string(data)}, " ")
}
