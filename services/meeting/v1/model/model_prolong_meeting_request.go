package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ProlongMeetingRequest struct {

	// 会议ID
	ConferenceID string `json:"conferenceID" xml:"conferenceID"`

	// 会控授权令牌，通过获取会控token接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization" xml:"X-Conference-Authorization"`

	Body *RestProlongDurReqBody `json:"body,omitempty" xml:"body"`
}

func (o ProlongMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProlongMeetingRequest struct{}"
	}

	return strings.Join([]string{"ProlongMeetingRequest", string(data)}, " ")
}
