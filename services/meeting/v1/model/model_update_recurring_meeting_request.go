package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRecurringMeetingRequest struct {

	// 标识是否为第三方portal过来的请求
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty" xml:"X-Authorization-Type"`

	// 用户的uuid（已在USG注册过的）
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 用于区分到哪个HCSO站点鉴权。
	XSiteId *string `json:"X-Site-Id,omitempty" xml:"X-Site-Id"`

	// 会议标识
	ConferenceID string `json:"conferenceID" xml:"conferenceID"`

	Body *RestScheduleConfDto `json:"body,omitempty" xml:"body"`
}

func (o UpdateRecurringMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecurringMeetingRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecurringMeetingRequest", string(data)}, " ")
}
