package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityReportSubscriptionsResponse Response Object
type ListSecurityReportSubscriptionsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Items          *[]SecurityReportSubscriptionSummaryResponse `json:"items,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o ListSecurityReportSubscriptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityReportSubscriptionsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityReportSubscriptionsResponse", string(data)}, " ")
}
