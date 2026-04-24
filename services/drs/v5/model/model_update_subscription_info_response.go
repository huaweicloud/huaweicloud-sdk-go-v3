package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubscriptionInfoResponse Response Object
type UpdateSubscriptionInfoResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateSubscriptionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionInfoResponse", string(data)}, " ")
}
