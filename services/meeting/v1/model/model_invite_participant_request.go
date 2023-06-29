package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InviteParticipantRequest Request Object
type InviteParticipantRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
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
