package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityReportSubscriptionRequest Request Object
type ShowSecurityReportSubscriptionRequest struct {

	// 安全报告订阅ID
	SubscriptionId string `json:"subscription_id"`
}

func (o ShowSecurityReportSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityReportSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityReportSubscriptionRequest", string(data)}, " ")
}
