package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionOrderResponse Response Object
type DeleteSubscriptionOrderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubscriptionOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionOrderResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionOrderResponse", string(data)}, " ")
}
