package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMeetingRequest Request Object
type UpdateMeetingRequest struct {

	// 会议ID。 > 创建会议时返回的conferenceID。不是vmrConferenceID。
	ConferenceID string `json:"conferenceID"`

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
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
