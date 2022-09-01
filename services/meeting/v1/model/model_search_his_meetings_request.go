package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchHisMeetingsRequest struct {

	// 用户UUID。 管理员有权限查询权限范围内的所有帐号，普通帐号仅能查询自己的。
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 指定返回的与会者列表的记录索引。该值必须大于等于0； 默认为0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 指定返回的记录数。默认值为20，最大值为500。 当pageSize大于最大值500时，系统会默认设置为500。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 根据会议主题，预定人和会议id关键词的字符串，查询历史会议。
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey"`

	// 指定是否查询企业下所有用户的会议记录，如果登录帐号不是企业管理员，则该字段无效。如果该字段为true，则userUUID字段无效。 default: false
	QueryAll *bool `json:"queryAll,omitempty" xml:"queryAll"`

	// 查询的起始日期毫秒数。例如：1583078400000
	StartDate int64 `json:"startDate" xml:"startDate"`

	// 查询的截止日期毫秒数。例如：1585756799000
	EndDate int64 `json:"endDate" xml:"endDate"`

	// - ASC_StartTIME：根据会议开始时间升序排序。 - DSC_StartTIME：根据会议开始时间降序排序。 - ASC_RecordTYPE：根据是否具有录播文件排序，之后默认按照会议开始时间升序排序。 - DSC_RecordTYPE：根据是否含有录播文件排序，之后默认按照会议开始时间降序排序。
	SortType *string `json:"sortType,omitempty" xml:"sortType"`

	// 标识是否为第三方portal过来的请求。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty" xml:"X-Authorization-Type"`

	// 用于区分到哪个HCSO站点鉴权。
	XSiteId *string `json:"X-Site-Id,omitempty" xml:"X-Site-Id"`
}

func (o SearchHisMeetingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchHisMeetingsRequest struct{}"
	}

	return strings.Join([]string{"SearchHisMeetingsRequest", string(data)}, " ")
}
