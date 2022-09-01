package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelRecurringMeetingRequest struct {

	// 标识是否为第三方portal过来的请求
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty" xml:"X-Authorization-Type"`

	// 用户的uuid（已在USG注册过的）
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 用于区分到哪个HCSO站点鉴权。
	XSiteId *string `json:"X-Site-Id,omitempty" xml:"X-Site-Id"`

	// 会议标识
	ConferenceID string `json:"conferenceID" xml:"conferenceID"`

	// 取消会议操作类型,1表示需要结束在线会议
	Type *int32 `json:"type,omitempty" xml:"type"`
}

func (o CancelRecurringMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRecurringMeetingRequest struct{}"
	}

	return strings.Join([]string{"CancelRecurringMeetingRequest", string(data)}, " ")
}
