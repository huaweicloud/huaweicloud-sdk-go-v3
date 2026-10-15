package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetSubscriptionResponse Response Object
type ResetSubscriptionResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ResetSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ResetSubscriptionResponse", string(data)}, " ")
}
