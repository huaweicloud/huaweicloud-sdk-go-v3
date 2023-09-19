package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartLiveRoomsResponse Response Object
type ListSmartLiveRoomsResponse struct {

	// 直播间总数。
	Count *int32 `json:"count,omitempty"`

	// 直播间列表。
	SmartLiveRooms *[]SmartLiveRoomBaseInfo `json:"smart_live_rooms,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSmartLiveRoomsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartLiveRoomsResponse struct{}"
	}

	return strings.Join([]string{"ListSmartLiveRoomsResponse", string(data)}, " ")
}
