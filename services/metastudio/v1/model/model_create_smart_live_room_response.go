package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSmartLiveRoomResponse Response Object
type CreateSmartLiveRoomResponse struct {

	// 直播间ID
	RoomId *string `json:"room_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSmartLiveRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSmartLiveRoomResponse struct{}"
	}

	return strings.Join([]string{"CreateSmartLiveRoomResponse", string(data)}, " ")
}
