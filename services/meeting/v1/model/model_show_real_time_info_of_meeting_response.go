package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRealTimeInfoOfMeetingResponse struct {
	// 所有参加会议的与会者列表，包括未入会的以及在线的与会者信息。

	Attendees *[]RealTimeAttendee `json:"attendees,omitempty"`
	// 在线会场列表，包括已进入会议、呼叫中、正在加入会议的与会者列表等。

	Participants *[]RealTimeParticipant `json:"participants,omitempty"`

	ConfInfo       *RealTimeConfInfo `json:"confInfo,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowRealTimeInfoOfMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealTimeInfoOfMeetingResponse struct{}"
	}

	return strings.Join([]string{"ShowRealTimeInfoOfMeetingResponse", string(data)}, " ")
}
