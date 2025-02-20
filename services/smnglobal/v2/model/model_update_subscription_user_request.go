package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubscriptionUserRequest Request Object
type UpdateSubscriptionUserRequest struct {

	// 订阅用户ID。
	Id string `json:"id"`

	Body *UpdateSubscriptionUserRequestBody `json:"body,omitempty"`
}

func (o UpdateSubscriptionUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionUserRequest", string(data)}, " ")
}
