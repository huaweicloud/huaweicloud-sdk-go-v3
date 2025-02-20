package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionUserResponse Response Object
type CreateSubscriptionUserResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 订阅用户ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSubscriptionUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserResponse struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserResponse", string(data)}, " ")
}
