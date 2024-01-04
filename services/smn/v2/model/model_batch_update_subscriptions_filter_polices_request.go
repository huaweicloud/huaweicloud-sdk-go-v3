package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateSubscriptionsFilterPolicesRequest Request Object
type BatchUpdateSubscriptionsFilterPolicesRequest struct {
	Body *BatchSubscriptionsFilterPolicesRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateSubscriptionsFilterPolicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateSubscriptionsFilterPolicesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateSubscriptionsFilterPolicesRequest", string(data)}, " ")
}
