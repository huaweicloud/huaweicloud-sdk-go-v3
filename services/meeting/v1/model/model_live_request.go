package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type LiveRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
	// 会控授权令牌，通过获取会控token接口获得。

	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestSetLiveReqBody `json:"body,omitempty"`
}

func (o LiveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveRequest struct{}"
	}

	return strings.Join([]string{"LiveRequest", string(data)}, " ")
}
