package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRealTimeInfoOfMeetingRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 会控授权令牌，通过获取会控token接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`
}

func (o ShowRealTimeInfoOfMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealTimeInfoOfMeetingRequest struct{}"
	}

	return strings.Join([]string{"ShowRealTimeInfoOfMeetingRequest", string(data)}, " ")
}
