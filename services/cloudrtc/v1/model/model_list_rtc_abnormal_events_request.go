package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcAbnormalEventsRequest Request Object
type ListRtcAbnormalEventsRequest struct {

	// 应用ID
	App string `json:"app"`

	// 房间ID
	RoomId *string `json:"room_id,omitempty"`

	// 用户ID
	Uid *string `json:"uid,omitempty"`

	// 查询起始时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T06:00:00Z，不填写则默认读取过去1小时数据数据。
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T07:00:00Z，不填写则默认为当前时间。
	EndTime *string `json:"end_time,omitempty"`

	// 查询结果条数
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRtcAbnormalEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcAbnormalEventsRequest struct{}"
	}

	return strings.Join([]string{"ListRtcAbnormalEventsRequest", string(data)}, " ")
}
