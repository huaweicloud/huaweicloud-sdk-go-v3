package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityReportSubscriptionRequest Request Object
type CreateSecurityReportSubscriptionRequest struct {
	Body *CreateSecurityReportSubscriptionRequestBody `json:"body,omitempty"`
}

func (o CreateSecurityReportSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityReportSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityReportSubscriptionRequest", string(data)}, " ")
}
