package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionsByTopicResponse Response Object
type DeleteSubscriptionsByTopicResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSubscriptionsByTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionsByTopicResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionsByTopicResponse", string(data)}, " ")
}
