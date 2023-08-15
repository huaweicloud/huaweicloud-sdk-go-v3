package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetParticipantViewRequest Request Object
type SetParticipantViewRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 专业会议终端的与会者标识。
	ParticipantID string `json:"participantID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestParticipantViewReqBody `json:"body,omitempty"`
}

func (o SetParticipantViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetParticipantViewRequest struct{}"
	}

	return strings.Join([]string{"SetParticipantViewRequest", string(data)}, " ")
}
