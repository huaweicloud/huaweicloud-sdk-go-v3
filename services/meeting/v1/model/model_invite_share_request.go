package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type InviteShareRequest struct {
	// 会议id，创建会议时生成

	ConferenceID string `json:"conferenceID"`
	// 会场id,可以通过查询会场id接口获取

	ParticipantID string `json:"participantID"`
	// 会控授权令牌，通过调用申请会控token的接口生成

	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *InviteShareDto `json:"body,omitempty"`
}

func (o InviteShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteShareRequest struct{}"
	}

	return strings.Join([]string{"InviteShareRequest", string(data)}, " ")
}
