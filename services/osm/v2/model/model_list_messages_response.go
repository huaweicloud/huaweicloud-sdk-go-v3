package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMessagesResponse struct {
	// 总数

	Count *int32 `json:"count,omitempty"`
	// 留言列表

	MessageList    *[]QueryMessageInfoV2 `json:"message_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessagesResponse struct{}"
	}

	return strings.Join([]string{"ListMessagesResponse", string(data)}, " ")
}
