package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type InviteOperateVideoRequest struct {

	// 会议id，创建会议时生成
	ConferenceID string `json:"conferenceID"`

	// 被操作的会场id，可以通过查询会场id接口获取。
	ParticipantID string `json:"participantID"`

	// 会控授权令牌，通过调用申请会控token的接口生成
	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestVideoBody `json:"body,omitempty"`
}

func (o InviteOperateVideoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteOperateVideoRequest struct{}"
	}

	return strings.Join([]string{"InviteOperateVideoRequest", string(data)}, " ")
}
