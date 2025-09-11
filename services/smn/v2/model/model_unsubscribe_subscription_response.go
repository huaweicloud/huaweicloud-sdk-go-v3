package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeSubscriptionResponse Response Object
type UnsubscribeSubscriptionResponse struct {

	// 响应信息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnsubscribeSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"UnsubscribeSubscriptionResponse", string(data)}, " ")
}
