package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecurringMeetingRequest Request Object
type CreateRecurringMeetingRequest struct {

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
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
