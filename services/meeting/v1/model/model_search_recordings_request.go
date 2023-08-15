package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchRecordingsRequest Request Object
type SearchRecordingsRequest struct {

	// 用户的UUID。 > 仅管理员有权限查询本企业其他用户的会议录制；普通帐号该字段无效，只能查询自己的。
	UserUUID *string `json:"userUUID,omitempty"`

	// 指定返回的页面索引。该值必须大于0。 default: 0
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。默认是20，最大500条。
	Limit *int32 `json:"limit,omitempty"`

	// 是否查询企业下所有用户的历史会议。 * true：查询所有用户的会议录制 * false：仅查询管理员自己的会议录制 > 仅对企业管理员生效。
	QueryAll *bool `json:"queryAll,omitempty"`

	// 查询条件。会议主题、会议预约人和会议ID等可作为搜索内容。
	SearchKey *string `json:"searchKey,omitempty"`

	// 查询的起始时间戳（单位毫秒）。
	StartDate int64 `json:"startDate"`

	// 查询的截止时间戳（单位毫秒）。
	EndDate int64 `json:"endDate"`

	// 查询结果排序类型。 - ASC_StartTIME：按录制开始时间升序排序 - DSC_StartTIME：按录制开始时间降序排序
	SortType *string `json:"sortType,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o SearchRecordingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchRecordingsRequest struct{}"
	}

	return strings.Join([]string{"SearchRecordingsRequest", string(data)}, " ")
}
