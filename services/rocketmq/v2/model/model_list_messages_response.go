package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMessagesResponse struct {

	// 消息列表。
	Messages *[]Message `json:"messages,omitempty"`

	// 消息总数。
	Total          float32 `json:"total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessagesResponse struct{}"
	}

	return strings.Join([]string{"ListMessagesResponse", string(data)}, " ")
}
