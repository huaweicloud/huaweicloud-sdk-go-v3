package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetRoleRequest struct {
	// 会议ID

	ConferenceID string `json:"conferenceID"`
	// 与会者标识。

	ParticipantID string `json:"participantID"`
	// 会控授权令牌，通过获取会控token接口获得。

	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestChairTokenReqBody `json:"body,omitempty"`
}

func (o SetRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRoleRequest struct{}"
	}

	return strings.Join([]string{"SetRoleRequest", string(data)}, " ")
}
