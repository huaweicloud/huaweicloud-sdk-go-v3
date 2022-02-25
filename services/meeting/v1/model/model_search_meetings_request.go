package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchMeetingsRequest struct {
	// 待查询的会议预定者的uuid 仅管理员有权限查询权限范围内的所有账号；普通账号该字段无效，只能查询自己的。

	UserUUID *string `json:"userUUID,omitempty"`
	// 指定返回的记录索引。该值必须大于等于0； 默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 指定返回的记录数，默认值由会议AS定义，默认是20，最大500条。

	Limit *int32 `json:"limit,omitempty"`
	// 指定是否查询企业下所有用户的会议记录。 如果登录帐号不是企业管理员，则该字段无效。 如果该字段为true，则userUUID字段无效。 default : false

	QueryAll *bool `json:"queryAll,omitempty"`
	// 查询用来当作关键词的字符串。长度限制为1-128个字符。

	SearchKey *string `json:"searchKey,omitempty"`
	// - ADAY:  一天。 - AWEEK:  一周。 - AMONTH:  一个月。 - ALL:  查询所有。

	QueryConfMode *string `json:"queryConfMode,omitempty"`
	// - ASC_StartTIME:  按会议开始时间升序排序。 - DSC_StartTIME:  按会议开始时间降序排序。

	SortType *string `json:"sortType,omitempty"`
	// 标识是否为第三方portal过来的请求。

	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`
	// 用于区分到哪个HCSO站点鉴权。

	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o SearchMeetingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchMeetingsRequest struct{}"
	}

	return strings.Join([]string{"SearchMeetingsRequest", string(data)}, " ")
}
