package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateWebSocketTokenRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
	// 会控正式Token。 该头域统一为BASE64编码。

	XConferenceAuthorization string `json:"X-Conference-Authorization"`
}

func (o CreateWebSocketTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWebSocketTokenRequest struct{}"
	}

	return strings.Join([]string{"CreateWebSocketTokenRequest", string(data)}, " ")
}
