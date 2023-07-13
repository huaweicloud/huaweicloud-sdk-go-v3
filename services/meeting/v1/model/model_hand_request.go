package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandRequest Request Object
type HandRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 与会者标识。
	ParticipantID string `json:"participantID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestHandsUpReqBody `json:"body,omitempty"`
}

func (o HandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandRequest struct{}"
	}

	return strings.Join([]string{"HandRequest", string(data)}, " ")
}
