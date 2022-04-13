package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RollcallParticipantRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
	// 与会者标识。

	ParticipantID string `json:"participantID"`
	// 会控授权令牌，通过获取会控token接口获得。

	XConferenceAuthorization string `json:"X-Conference-Authorization"`
}

func (o RollcallParticipantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollcallParticipantRequest struct{}"
	}

	return strings.Join([]string{"RollcallParticipantRequest", string(data)}, " ")
}
