package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartChatRoomsResponse Response Object
type ListSmartChatRoomsResponse struct {

	// 智能交互对话直播间总数。
	Count *int32 `json:"count,omitempty"`

	// 智能交互对话直播间列表。
	SmartChatRooms *[]SmartChatRoomBaseInfo `json:"smart_chat_rooms,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSmartChatRoomsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartChatRoomsResponse struct{}"
	}

	return strings.Join([]string{"ListSmartChatRoomsResponse", string(data)}, " ")
}
