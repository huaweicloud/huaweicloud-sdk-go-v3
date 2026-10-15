package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartSubscriptionResponse Response Object
type StartSubscriptionResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o StartSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"StartSubscriptionResponse", string(data)}, " ")
}
