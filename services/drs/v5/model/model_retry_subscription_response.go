package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetrySubscriptionResponse Response Object
type RetrySubscriptionResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o RetrySubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetrySubscriptionResponse struct{}"
	}

	return strings.Join([]string{"RetrySubscriptionResponse", string(data)}, " ")
}
