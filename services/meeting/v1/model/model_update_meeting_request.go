package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMeetingRequest struct {
	// 会议ID

	ConferenceID string `json:"conferenceID"`
	// 用户的UUID（已在USG注册过的）。

	UserUUID *string `json:"userUUID,omitempty"`
	// 标识是否为第三方portal过来的请求。

	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`
	// 用于区分到哪个HCSO站点鉴权。

	XSiteId *string `json:"X-Site-Id,omitempty"`

	Body *RestScheduleConfDto `json:"body,omitempty"`
}

func (o UpdateMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMeetingRequest struct{}"
	}

	return strings.Join([]string{"UpdateMeetingRequest", string(data)}, " ")
}
