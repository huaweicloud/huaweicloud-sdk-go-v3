package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRecurringSubMeetingRequest Request Object
type CancelRecurringSubMeetingRequest struct {

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 取消会议操作类型。默认已召开的会议不能取消。 * 1：需要结束正在召开的会议
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
