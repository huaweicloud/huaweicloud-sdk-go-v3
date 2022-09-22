package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRealTimeInfoOfMeetingResponse struct {

	// 被邀请与会者信息，包括预约会议时邀请的与会者和会中主持人邀请的与会者，已经加入会议的和未加入会议的都返回。
	Attendees *[]RealTimeAttendee `json:"attendees,omitempty"`

	// 在线与会者列表信息，包括已加入会议、被邀请正在呼叫中、正在加入会议的与会者列表等。 > * 同一个帐号用不同类型终端（手机端或者PC端等）加入会议时，是不同的在线与会者 > * 未加入或者已离会与会者，不在在线与会者列表中
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
