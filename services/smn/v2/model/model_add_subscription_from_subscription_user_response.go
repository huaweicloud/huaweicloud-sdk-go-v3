package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSubscriptionFromSubscriptionUserResponse Response Object
type AddSubscriptionFromSubscriptionUserResponse struct {

	// 添加订阅返回结果。
	SubscriptionsResult *[]AddSubscriptionFromSubscriptionUserResponseItem `json:"subscriptions_result,omitempty"`
	HttpStatusCode      int                                                `json:"-"`
}

func (o AddSubscriptionFromSubscriptionUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionFromSubscriptionUserResponse struct{}"
	}

	return strings.Join([]string{"AddSubscriptionFromSubscriptionUserResponse", string(data)}, " ")
}
