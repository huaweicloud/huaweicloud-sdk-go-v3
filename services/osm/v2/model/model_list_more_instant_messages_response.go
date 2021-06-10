package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListMoreInstantMessagesResponse struct {
	// 总数

	Count *int32 `json:"count,omitempty"`
	// 留言列表

	MessageList    *[]QueryMessageInfoV2 `json:"message_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListMoreInstantMessagesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMoreInstantMessagesResponse struct{}"
	}

	return strings.Join([]string{"ListMoreInstantMessagesResponse", string(data)}, " ")
}
