package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubscriptionsRequest Request Object
type BatchDeleteSubscriptionsRequest struct {
	Body *BatchDeleteSubscriptionsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSubscriptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubscriptionsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubscriptionsRequest", string(data)}, " ")
}
