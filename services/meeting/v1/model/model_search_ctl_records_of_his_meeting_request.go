package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCtlRecordsOfHisMeetingRequest Request Object
type SearchCtlRecordsOfHisMeetingRequest struct {

	// 会议UUID。
	ConfUUID string `json:"confUUID"`

	// 查询偏移量。默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。默认值20，最大500条。
	Limit *int32 `json:"limit,omitempty"`

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`

	// 语言。默认简体中文。 - zh-CN: 简体中文。 - en-US: 美国英文。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`
}

func (o SearchCtlRecordsOfHisMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCtlRecordsOfHisMeetingRequest struct{}"
	}

	return strings.Join([]string{"SearchCtlRecordsOfHisMeetingRequest", string(data)}, " ")
}
