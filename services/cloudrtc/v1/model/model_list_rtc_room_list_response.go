package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcRoomListResponse Response Object
type ListRtcRoomListResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 查询结果限制
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 房间列表信息
	RoomInfoList *[]RtcServerRoomInfo `json:"room_info_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcRoomListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcRoomListResponse struct{}"
	}

	return strings.Join([]string{"ListRtcRoomListResponse", string(data)}, " ")
}
