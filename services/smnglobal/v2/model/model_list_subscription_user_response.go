package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionUserResponse Response Object
type ListSubscriptionUserResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 订阅用户数量。
	Count *int32 `json:"count,omitempty"`

	// 订阅用户信息列表。
	SubscriptionUsers *[]ListSubscriptionUserResponseItemInfo `json:"subscription_users,omitempty"`
	HttpStatusCode    int                                     `json:"-"`
}

func (o ListSubscriptionUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserResponse", string(data)}, " ")
}
