package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRtcRealtimeScaleDimensionRequest struct {

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

	// 查询的数据类型  OnlineUsers:在线用户数
	Metric string `json:"metric"`

	// 维度类型: region:省份 access_net:网络类型 platform:系统平台 sdk:SDK版本
	Dimension string `json:"dimension"`

	// 查询时刻。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ
	Time string `json:"time"`
}

func (o ListRtcRealtimeScaleDimensionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcRealtimeScaleDimensionRequest struct{}"
	}

	return strings.Join([]string{"ListRtcRealtimeScaleDimensionRequest", string(data)}, " ")
}
