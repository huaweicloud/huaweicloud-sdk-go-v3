package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRecurringMeetingRequest struct {

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	Body *RestScheduleConfDto `json:"body,omitempty"`
}

func (o UpdateRecurringMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecurringMeetingRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecurringMeetingRequest", string(data)}, " ")
}
