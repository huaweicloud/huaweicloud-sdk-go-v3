package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockViewRequest Request Object
type LockViewRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 专业会议终端的与会者标识。
	ParticipantID string `json:"participantID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestLockSiteViewReqBody `json:"body,omitempty"`
}

func (o LockViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockViewRequest struct{}"
	}

	return strings.Join([]string{"LockViewRequest", string(data)}, " ")
}
