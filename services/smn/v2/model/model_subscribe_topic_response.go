package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeTopicResponse Response Object
type SubscribeTopicResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 订阅者的唯一的资源标识。
	SubscriptionUrn *string `json:"subscription_urn,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o SubscribeTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeTopicResponse struct{}"
	}

	return strings.Join([]string{"SubscribeTopicResponse", string(data)}, " ")
}
