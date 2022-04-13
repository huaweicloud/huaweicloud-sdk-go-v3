package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddSubscriptionResponse struct {
	// 请求的唯一标识ID。

	RequestId *string `json:"request_id,omitempty"`
	// 订阅者的唯一资源标识。

	SubscriptionUrn *string `json:"subscription_urn,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o AddSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"AddSubscriptionResponse", string(data)}, " ")
}
