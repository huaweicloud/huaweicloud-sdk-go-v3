package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecordingsRequest Request Object
type DeleteRecordingsRequest struct {

	// 会议UUID列表，多个会议UUID之间以英文逗号隔开。
	ConfUUIDs string `json:"confUUIDs"`

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o DeleteRecordingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordingsRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordingsRequest", string(data)}, " ")
}
