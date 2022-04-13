package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRecurringSubMeetingRequest struct {
	// 标识是否为第三方portal过来的请求

	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`
	// 用户的uuid（已在USG注册过的）

	UserUUID *string `json:"userUUID,omitempty"`
	// 用于区分到哪个HCSO站点鉴权。

	XSiteId *string `json:"X-Site-Id,omitempty"`
	// 会议标识

	ConferenceID string `json:"conferenceID"`

	Body *RestScheduleCycleConfDto `json:"body,omitempty"`
}

func (o UpdateRecurringSubMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecurringSubMeetingRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecurringSubMeetingRequest", string(data)}, " ")
}
