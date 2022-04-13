package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelRecurringSubMeetingRequest struct {
	// 标识是否为第三方portal过来的请求

	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`
	// 用户的uuid（已在USG注册过的）

	UserUUID *string `json:"userUUID,omitempty"`
	// 用于区分到哪个HCSO站点鉴权。

	XSiteId *string `json:"X-Site-Id,omitempty"`
	// 会议标识

	ConferenceID string `json:"conferenceID"`
	// 取消会议操作类型,1表示需要结束在线会议

	Type *int32 `json:"type,omitempty"`

	Body *RestCancelSingleRecordCycleConfListReqBody `json:"body,omitempty"`
}

func (o CancelRecurringSubMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRecurringSubMeetingRequest struct{}"
	}

	return strings.Join([]string{"CancelRecurringSubMeetingRequest", string(data)}, " ")
}
