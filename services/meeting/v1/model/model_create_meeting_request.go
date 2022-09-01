package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMeetingRequest struct {

	// 用户的UUID（已在USG注册过的）。
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 标识是否为第三方portal过来的请求。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty" xml:"X-Authorization-Type"`

	// 用于区分到哪个HCSO站点鉴权。
	XSiteId *string `json:"X-Site-Id,omitempty" xml:"X-Site-Id"`

	Body *RestScheduleConfDto `json:"body,omitempty" xml:"body"`
}

func (o CreateMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMeetingRequest struct{}"
	}

	return strings.Join([]string{"CreateMeetingRequest", string(data)}, " ")
}
