package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubscriptionUserResponse Response Object
type UpdateSubscriptionUserResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSubscriptionUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionUserResponse", string(data)}, " ")
}
