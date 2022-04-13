package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRecurringMeetingRequest struct {
	// 标识是否为第三方portal过来的请求

	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`
	// 用户的uuid（已在USG注册过的）

	UserUUID *string `json:"userUUID,omitempty"`
	// 用于区分到哪个HCSO站点鉴权。

	XSiteId *string `json:"X-Site-Id,omitempty"`

	Body *RestScheduleConfDto `json:"body,omitempty"`
}

func (o CreateRecurringMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecurringMeetingRequest struct{}"
	}

	return strings.Join([]string{"CreateRecurringMeetingRequest", string(data)}, " ")
}
