package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealTimeInfoOfMeetingRequest Request Object
type ShowRealTimeInfoOfMeetingRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`
}

func (o ShowRealTimeInfoOfMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealTimeInfoOfMeetingRequest struct{}"
	}

	return strings.Join([]string{"ShowRealTimeInfoOfMeetingRequest", string(data)}, " ")
}
