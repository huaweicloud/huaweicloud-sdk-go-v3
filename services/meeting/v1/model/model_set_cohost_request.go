package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetCohostRequest Request Object
type SetCohostRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 与会者标识。
	ParticipantID string `json:"participantID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestSetCohostBody `json:"body,omitempty"`
}

func (o SetCohostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetCohostRequest struct{}"
	}

	return strings.Join([]string{"SetCohostRequest", string(data)}, " ")
}
