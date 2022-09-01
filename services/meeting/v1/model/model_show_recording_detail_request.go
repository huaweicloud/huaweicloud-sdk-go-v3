package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRecordingDetailRequest struct {

	// 会议的ConfUUID(通过查询录制列表获取)。
	ConfUUID string `json:"confUUID" xml:"confUUID"`

	// 用户的UUID（已在USG注册过的）。
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 标识是否为第三方portal过来的请求。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty" xml:"X-Authorization-Type"`

	// 用于区分到哪个HCSO站点鉴权。
	XSiteId *string `json:"X-Site-Id,omitempty" xml:"X-Site-Id"`
}

func (o ShowRecordingDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordingDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordingDetailRequest", string(data)}, " ")
}
