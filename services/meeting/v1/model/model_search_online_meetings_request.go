package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchOnlineMeetingsRequest Request Object
type SearchOnlineMeetingsRequest struct {

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 查询偏移量。默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。默认是20，最大500条。
	Limit *int32 `json:"limit,omitempty"`

	// 指定是否查询企业下所有用户的在线会议。默认值是false。 * true：查询所有用户的在线会议 * false：仅查询管理员自己的在线会议 > 仅对企业管理员生效。
	QueryAll *bool `json:"queryAll,omitempty"`

	// 查询条件 。会议主题、会议预约人和会议ID等可作为搜索内容。长度限制为1-128个字符。
	SearchKey *string `json:"searchKey,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o SearchOnlineMeetingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchOnlineMeetingsRequest struct{}"
	}

	return strings.Join([]string{"SearchOnlineMeetingsRequest", string(data)}, " ")
}
