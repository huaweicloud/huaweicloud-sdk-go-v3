package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowClientRecordRequest Request Object
type AllowClientRecordRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 会控Token，通过[[获取会控token](https://support.huaweicloud.com/api-meeting/meeting_21_0027.html)](tag:hws)[[获取会控token](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0027.html)](tag:hk)接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	// 与会者标识。
	ParticipantID string `json:"participantID"`

	Body *RestAllowClientRecordReqBody `json:"body,omitempty"`
}

func (o AllowClientRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowClientRecordRequest struct{}"
	}

	return strings.Join([]string{"AllowClientRecordRequest", string(data)}, " ")
}
