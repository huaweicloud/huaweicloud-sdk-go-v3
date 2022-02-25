package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchRecordingsRequest struct {
	// 待查询的会议预定者的用户UUID。仅管理员有权限查询权限范围内的所有录制，普通帐号仅能查询自己的。 默认是登录帐号。

	UserUUID *string `json:"userUUID,omitempty"`
	// 指定返回的页面索引。该值必须大于0。 default: 0

	Offset *int32 `json:"offset,omitempty"`
	// 指定返回的记录数。默认值为20，最大值为100。

	Limit *int32 `json:"limit,omitempty"`
	// 指定是否查询企业下所有用户的会议录制。 - 如果登录帐号不是企业管理员，则该字段无效。 - 如果该字段为true，则userUUID字段无效。

	QueryAll *bool `json:"queryAll,omitempty"`
	// 会议主题，预定人或会议id可作为搜索词，查询录制。

	SearchKey *string `json:"searchKey,omitempty"`
	// 查询的起始日期毫秒数。

	StartDate int64 `json:"startDate"`
	// 查询的截止日期毫秒数。

	EndDate int64 `json:"endDate"`
	// - ASC_StartTIME：按录制开始时间升序排序。 - DSC_StartTIME：按录制开始时间降序排序。

	SortType *string `json:"sortType,omitempty"`
	// 标识是否为第三方portal过来的请求。

	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`
	// 用于区分到哪个HCSO站点鉴权。

	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o SearchRecordingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchRecordingsRequest struct{}"
	}

	return strings.Join([]string{"SearchRecordingsRequest", string(data)}, " ")
}
