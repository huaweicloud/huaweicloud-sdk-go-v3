package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmSubscriptionResponse Response Object
type ConfirmSubscriptionResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 订阅者的唯一的资源标识。
	SubscriptionUrn *string `json:"subscription_urn,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ConfirmSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ConfirmSubscriptionResponse", string(data)}, " ")
}
