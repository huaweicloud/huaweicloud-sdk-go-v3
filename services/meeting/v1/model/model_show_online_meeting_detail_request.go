package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOnlineMeetingDetailRequest Request Object
type ShowOnlineMeetingDetailRequest struct {

	// 会议ID。 > 会议ID。创建会议时返回的conferenceID。不是vmrConferenceID。
	ConferenceID string `json:"conferenceID"`

	// 查询偏移量。默认为0。针对PageParticipant中的与会者分页。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。默认值20。
	Limit *int32 `json:"limit,omitempty"`

	// 查询条件。长度限制为1-128个字符。
	SearchKey *string `json:"searchKey,omitempty"`

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 默认值为0。 - 0: 不区分终端和与会人 - 1: 分页查询区分终端和与会人，结果合并返回 - 2: 单独查询终端和与会人，结果单独返回
	XType *string `json:"X-Type,omitempty"`

	// 当X-Type为2时，该字段有效。默认值为0。 - 0: 查询与会人 - 1: 查询终端
	XQueryType *string `json:"X-Query-Type,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o ShowOnlineMeetingDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOnlineMeetingDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowOnlineMeetingDetailRequest", string(data)}, " ")
}
