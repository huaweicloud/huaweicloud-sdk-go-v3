package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartChatRoomsResponse Response Object
type ListSmartChatRoomsResponse struct {

	// **参数解释**： 智能交互对话总数。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 智能交互对话总并发路数。
	CountConcurrency *int32 `json:"count_concurrency,omitempty"`

	// **参数解释**： 智能交互对话按接入端计费总路数。
	CountClientNums *int32 `json:"count_client_nums,omitempty"`

	// **参数解释**： 智能交互对话按接入端计费（Token）总路数。
	CountClientNumsToken *int32 `json:"count_client_nums_token,omitempty"`

	// 智能交互对话列表。
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
