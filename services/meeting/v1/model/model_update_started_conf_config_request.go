package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateStartedConfConfigRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
	// 会控正式Token。 该头域统一为BASE64编码。

	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *UpdateStartedConfigReqBody `json:"body,omitempty"`
}

func (o UpdateStartedConfConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStartedConfConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateStartedConfConfigRequest", string(data)}, " ")
}
