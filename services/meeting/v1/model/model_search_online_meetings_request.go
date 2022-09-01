package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchOnlineMeetingsRequest struct {

	// 用户UUID。 仅管理员有权限查询权限范围内的所有帐号，普通帐号仅能查询自己的。
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 指定返回的记录索引。该值必须大于等于0； 默认为0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 指定返回的记录数。默认是20，最大值为500。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 指定是否查询企业下所有用户的会议记录。 如果登录帐号不是企业管理员，则该字段无效。 如果该字段为true，则userUUID字段无效。 default : false
	QueryAll *bool `json:"queryAll,omitempty" xml:"queryAll"`

	// 查询用来当作关键词的字符串。会议主题、会议预约人和会议ID等可作为搜索内容。
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey"`

	// 标识是否为第三方portal过来的请求。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty" xml:"X-Authorization-Type"`

	// 用于区分到哪个HCSO站点鉴权。
	XSiteId *string `json:"X-Site-Id,omitempty" xml:"X-Site-Id"`
}

func (o SearchOnlineMeetingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchOnlineMeetingsRequest struct{}"
	}

	return strings.Join([]string{"SearchOnlineMeetingsRequest", string(data)}, " ")
}
