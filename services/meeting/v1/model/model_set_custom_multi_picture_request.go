package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetCustomMultiPictureRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
	// 会控授权令牌，通过获取会控token接口获得。

	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestCustomMultiPictureBody `json:"body,omitempty"`
}

func (o SetCustomMultiPictureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetCustomMultiPictureRequest struct{}"
	}

	return strings.Join([]string{"SetCustomMultiPictureRequest", string(data)}, " ")
}
