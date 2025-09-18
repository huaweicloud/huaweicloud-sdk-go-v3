package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityReportSubscriptionRequest Request Object
type DeleteSecurityReportSubscriptionRequest struct {
	SubscriptionId string `json:"subscription_id"`
}

func (o DeleteSecurityReportSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityReportSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityReportSubscriptionRequest", string(data)}, " ")
}
