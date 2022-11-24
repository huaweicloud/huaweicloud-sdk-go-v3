package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowHisMeetingDetailRequest struct {

	// 会议UUID。
	ConfUUID string `json:"confUUID"`

	// 查询偏移量。默认为0。针对PageParticipant 中的与会者分页。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。默认值20。
	Limit *int32 `json:"limit,omitempty"`

	// 查询条件 。会议主题、会议预约人和会议ID等可作为搜索内容。长度限制为1-128个字符。
	SearchKey *string `json:"searchKey,omitempty"`

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 默认值为0。 0: 不区分会议室和与会人。 1：分页查询区分会议室和与会人，结果合并返回。 2：单独查询会议室与与会人，结果也是单独返回。
	XType *int32 `json:"X-Type,omitempty"`

	// 当X-Type为2时，该字段有效。默认值为0。 0: 查询与会人。 1：查询终端。
	XQueryType *int32 `json:"X-Query-Type,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o ShowHisMeetingDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHisMeetingDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowHisMeetingDetailRequest", string(data)}, " ")
}
