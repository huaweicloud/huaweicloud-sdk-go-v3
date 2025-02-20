package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionUserResponse Response Object
type DeleteSubscriptionUserResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSubscriptionUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionUserResponse", string(data)}, " ")
}
