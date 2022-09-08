package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type LockViewRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 与会者标识。
	ParticipantID string `json:"participantID"`

	// 会控授权令牌，通过获取会控token接口获得。
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
