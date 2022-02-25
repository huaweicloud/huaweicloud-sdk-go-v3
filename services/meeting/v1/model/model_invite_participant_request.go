package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type InviteParticipantRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
	// 会控授权令牌，通过获取会控token接口获得。

	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestInviteReqBody `json:"body,omitempty"`
}

func (o InviteParticipantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteParticipantRequest struct{}"
	}

	return strings.Join([]string{"InviteParticipantRequest", string(data)}, " ")
}
