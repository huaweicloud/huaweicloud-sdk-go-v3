package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcRealtimeScaleRequest Request Object
type ListRtcRealtimeScaleRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息，与路径参数中的项目ID相同。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 应用标识
	App string `json:"app"`

	// 房间ID
	RoomId *string `json:"room_id,omitempty"`

	// 查询的数据类型  OnlineUsers:在线用户数;  RoomCount:在线房间数;
	Metric string `json:"metric"`

	// 查询起始时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T06:00:00Z，不写默认读取过去1小时数据数据。
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T07:00:00Z，不写默认为当前时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListRtcRealtimeScaleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcRealtimeScaleRequest struct{}"
	}

	return strings.Join([]string{"ListRtcRealtimeScaleRequest", string(data)}, " ")
}
