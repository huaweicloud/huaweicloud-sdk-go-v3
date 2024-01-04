package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSubscriptionsFilterPolicesRequest Request Object
type BatchCreateSubscriptionsFilterPolicesRequest struct {
	Body *BatchSubscriptionsFilterPolicesRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSubscriptionsFilterPolicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubscriptionsFilterPolicesRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSubscriptionsFilterPolicesRequest", string(data)}, " ")
}
