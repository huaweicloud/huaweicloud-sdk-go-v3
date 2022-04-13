package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelMeetingRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
	// 用户的UUID（已在USG注册过的）

	UserUUID *string `json:"userUUID,omitempty"`
	// 取消会议操作类型，1表示需要结束在线会议。

	Type *int32 `json:"type,omitempty"`
	// 标识是否为第三方portal过来的请求。

	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`
	// 用于区分到哪个HCSO站点鉴权。

	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o CancelMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelMeetingRequest struct{}"
	}

	return strings.Join([]string{"CancelMeetingRequest", string(data)}, " ")
}
