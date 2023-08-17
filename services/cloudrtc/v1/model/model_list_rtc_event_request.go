package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcEventRequest Request Object
type ListRtcEventRequest struct {

	// 应用id
	AppId string `json:"app_id"`

	// 房间id
	RoomId *string `json:"room_id,omitempty"`

	// 起始时间。UTC时间，格式：yyyy-mm-ddThh:mm:ssZ，如2020-04-23T06:00:00Z。查询起止时间不超过1个小时，每次查询单个用户时，支持跨天查询，最长1天。
	StartTime string `json:"start_time"`

	// 结束时间。UTC时间，格式：yyyy-mm-ddThh:mm:ssZ，如2020-04-23T06:00:00Z。查询起止时间不超过1个小时，每次查询单个用户时，支持跨天查询，最长1天。
	EndTime string `json:"end_time"`
}

func (o ListRtcEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcEventRequest struct{}"
	}

	return strings.Join([]string{"ListRtcEventRequest", string(data)}, " ")
}
