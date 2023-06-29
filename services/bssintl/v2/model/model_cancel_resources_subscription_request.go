package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelResourcesSubscriptionRequest Request Object
type CancelResourcesSubscriptionRequest struct {
	Body *UnsubscribeResourcesReq `json:"body,omitempty"`
}

func (o CancelResourcesSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelResourcesSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"CancelResourcesSubscriptionRequest", string(data)}, " ")
}
