package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type HangUpRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID" xml:"conferenceID"`

	// 会控授权令牌，通过获取会控token接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization" xml:"X-Conference-Authorization"`

	Body *RestBulkHangUpReqBody `json:"body,omitempty" xml:"body"`
}

func (o HangUpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HangUpRequest struct{}"
	}

	return strings.Join([]string{"HangUpRequest", string(data)}, " ")
}
