package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListUnreadNewInstantMessagesResponse struct {
	// 未读消息列表

	Imunread       *[]ImUnreadV2 `json:"imunread,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListUnreadNewInstantMessagesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListUnreadNewInstantMessagesResponse struct{}"
	}

	return strings.Join([]string{"ListUnreadNewInstantMessagesResponse", string(data)}, " ")
}
