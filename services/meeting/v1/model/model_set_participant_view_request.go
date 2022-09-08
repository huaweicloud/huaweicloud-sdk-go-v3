package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetParticipantViewRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 与会者标识。
	ParticipantID string `json:"participantID"`

	// 会控授权令牌，通过获取会控token接口获得。
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
