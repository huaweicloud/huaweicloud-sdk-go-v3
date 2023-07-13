package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUnreadNewInstantMessagesResponse Response Object
type ListUnreadNewInstantMessagesResponse struct {

	// 未读消息列表
	Imunread       *[]ImUnreadV2 `json:"imunread,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListUnreadNewInstantMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUnreadNewInstantMessagesResponse struct{}"
	}

	return strings.Join([]string{"ListUnreadNewInstantMessagesResponse", string(data)}, " ")
}
