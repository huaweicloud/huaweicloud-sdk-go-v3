package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BroadcastParticipantRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
	// 与会者标识。

	ParticipantID string `json:"participantID"`
	// 会控授权令牌，通过获取会控token接口获得。

	XConferenceAuthorization string `json:"X-Conference-Authorization"`
}

func (o BroadcastParticipantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BroadcastParticipantRequest struct{}"
	}

	return strings.Join([]string{"BroadcastParticipantRequest", string(data)}, " ")
}
