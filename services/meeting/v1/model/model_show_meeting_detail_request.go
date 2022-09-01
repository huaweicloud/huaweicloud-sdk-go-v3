package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMeetingDetailRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID" xml:"conferenceID"`

	// 指定返回的与会者列表的记录索引。该值必须大于等于0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 指定返回的与会者记录数。默认值20。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 用来作关键词查询的字符串。长度限制为1-128个字符。
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey"`

	// 用户的UUID。第三方鉴权时需要携带。
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 默认值为0。 - 0: 不区分终端和与会人。 - 1: 分页查询区分终端和与会人，结果合并返回。 - 2: 单独查询终端和与会人，结果单独返回。
	XType *string `json:"X-Type,omitempty" xml:"X-Type"`

	// 当X-Type为2时，有效。默认为0。 - 0: 查询与会人。 - 1: 查询终端。
	XQueryType *string `json:"X-Query-Type,omitempty" xml:"X-Query-Type"`

	// 标识是否为第三方portal过来的请求。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty" xml:"X-Authorization-Type"`

	// 用于区分到哪个HCSO站点鉴权。
	XSiteId *string `json:"X-Site-Id,omitempty" xml:"X-Site-Id"`
}

func (o ShowMeetingDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeetingDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMeetingDetailRequest", string(data)}, " ")
}
