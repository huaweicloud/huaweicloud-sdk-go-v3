package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMeetingRequest struct {

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`

	Body *RestScheduleConfDto `json:"body,omitempty"`
}

func (o CreateMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMeetingRequest struct{}"
	}

	return strings.Join([]string{"CreateMeetingRequest", string(data)}, " ")
}
