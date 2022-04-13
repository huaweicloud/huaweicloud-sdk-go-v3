package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowOnlineMeetingDetailRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
	// 指定返回的记录索引。该值必须大于等于0； 默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 指定返回的记录数。

	Limit *int32 `json:"limit,omitempty"`
	// 用来作关键词查询的字符串。

	SearchKey *string `json:"searchKey,omitempty"`
	// 用户的UUID（已在USG注册过的）。

	UserUUID *string `json:"userUUID,omitempty"`
	// 默认值为0。 - 0: 不区分终端和与会人。 - 1: 分页查询区分终端和与会人，结果合并返回。 - 2: 单独查询终端和与会人，结果单独返回。

	XType *string `json:"X-Type,omitempty"`
	// 当X-Type为2时，该字段有效。默认值为0。 - 0: 查询与会人。 - 1: 查询终端。

	XQueryType *string `json:"X-Query-Type,omitempty"`
	// 标识是否为第三方portal过来的请求。

	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`
	// 用于区分到哪个HCSO站点鉴权。

	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o ShowOnlineMeetingDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOnlineMeetingDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowOnlineMeetingDetailRequest", string(data)}, " ")
}
