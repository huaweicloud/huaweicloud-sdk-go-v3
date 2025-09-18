package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityReportSubscriptionRequest Request Object
type UpdateSecurityReportSubscriptionRequest struct {
	SubscriptionId string `json:"subscription_id"`

	Body *UpdateSecurityReportSubscriptionRequestBody `json:"body,omitempty"`
}

func (o UpdateSecurityReportSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityReportSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityReportSubscriptionRequest", string(data)}, " ")
}
