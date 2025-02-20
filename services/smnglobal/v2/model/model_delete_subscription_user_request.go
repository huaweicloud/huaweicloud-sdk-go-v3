package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionUserRequest Request Object
type DeleteSubscriptionUserRequest struct {

	// 订阅用户ID。
	Id string `json:"id"`
}

func (o DeleteSubscriptionUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionUserRequest", string(data)}, " ")
}
