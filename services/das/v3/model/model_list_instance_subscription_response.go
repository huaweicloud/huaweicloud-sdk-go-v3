package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSubscriptionResponse Response Object
type ListInstanceSubscriptionResponse struct {
	Body           *[]ReportSubscription `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListInstanceSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceSubscriptionResponse", string(data)}, " ")
}
