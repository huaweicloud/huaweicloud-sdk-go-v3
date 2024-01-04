package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubscriptionsFilterPolicesRequest Request Object
type BatchDeleteSubscriptionsFilterPolicesRequest struct {
	Body *BatchDeleteSubscriptionsFilterPolicesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSubscriptionsFilterPolicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubscriptionsFilterPolicesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubscriptionsFilterPolicesRequest", string(data)}, " ")
}
