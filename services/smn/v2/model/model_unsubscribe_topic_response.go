package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeTopicResponse Response Object
type UnsubscribeTopicResponse struct {

	// 响应信息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnsubscribeTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeTopicResponse struct{}"
	}

	return strings.Join([]string{"UnsubscribeTopicResponse", string(data)}, " ")
}
