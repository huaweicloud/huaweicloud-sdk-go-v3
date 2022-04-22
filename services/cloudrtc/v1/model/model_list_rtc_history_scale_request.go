package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRtcHistoryScaleRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息，与路径参数中的项目ID相同。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 应用标识
	App string `json:"app"`

	// 查询的数据类型 - UserCount：通话人数，不同频道中的相同用户ID计为多人; - SessionCount：通话人次，用户每次加入频道计为一个通话人次; - RoomCount：房间数，从有用户加入房间到所有用户离开房间计为一个通话房间; - MaxOnlineUserCount：最大同时在线人数; - MaxOnlineRoomCount：最大同时在线房间数; - CommunicationDuration：音视频通话总时长; - VideoCommunicationDuration：视频通话总时长; - AudioCommunicationDuration：音频通话总时长;
	Metric []string `json:"metric"`

	// 查询起始时间。UTC时间，格式：YYYY-MM-DD，如2020-04-23，不写默认读取过去1天数据数据。
	StartDate *string `json:"start_date,omitempty"`

	// 查询结束时间。UTC时间，格式：YYYY-MM-DD，如2020-04-23
	EndDate *string `json:"end_date,omitempty"`
}

func (o ListRtcHistoryScaleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcHistoryScaleRequest struct{}"
	}

	return strings.Join([]string{"ListRtcHistoryScaleRequest", string(data)}, " ")
}
