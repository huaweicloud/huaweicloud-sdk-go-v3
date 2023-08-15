package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordingDetailRequest Request Object
type ShowRecordingDetailRequest struct {

	// 会议UUID(通过[[查询录制列表](https://support.huaweicloud.com/api-meeting/meeting_21_0048.html)](tag:hws)[[查询录制列表](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0048.html)](tag:hk)获取)。
	ConfUUID string `json:"confUUID"`

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o ShowRecordingDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordingDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordingDetailRequest", string(data)}, " ")
}
