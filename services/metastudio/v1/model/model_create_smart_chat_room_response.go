package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSmartChatRoomResponse Response Object
type CreateSmartChatRoomResponse struct {

	// 智能交互对话ID
	RoomId *string `json:"room_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSmartChatRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSmartChatRoomResponse struct{}"
	}

	return strings.Join([]string{"CreateSmartChatRoomResponse", string(data)}, " ")
}
