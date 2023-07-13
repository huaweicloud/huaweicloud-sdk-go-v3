package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchMeetingsRequest Request Object
type SearchMeetingsRequest struct {

	// 用户的UUID。 > 仅管理员有权限查询本企业其他用户的会议列表；普通帐号该字段无效，只能查询自己的。
	UserUUID *string `json:"userUUID,omitempty"`

	// 查询偏移量。默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。默认是20，最大500条。
	Limit *int32 `json:"limit,omitempty"`

	// 是否查询企业下所有用户的会议记录。默认值为false。 * true：查询所有用户的会议 * false：仅查询管理员自己创建的会议 > 仅对企业管理员生效。
	QueryAll *bool `json:"queryAll,omitempty"`

	// 查询条件。会议主题、会议预约人和会议ID等可作为搜索内容。长度限制为1-128个字符。
	SearchKey *string `json:"searchKey,omitempty"`

	// 查询时间范围。 - ADAY:  一天 - AWEEK:  一周 - AMONTH:  一个月 - ALL:  查询所有
	QueryConfMode *string `json:"queryConfMode,omitempty"`

	// 查询结果排序。 - ASC_StartTIME:  按会议开始时间升序排序 - DSC_StartTIME:  按会议开始时间降序排序
	SortType *string `json:"sortType,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o SearchMeetingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchMeetingsRequest struct{}"
	}

	return strings.Join([]string{"SearchMeetingsRequest", string(data)}, " ")
}
