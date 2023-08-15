package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchHisMeetingsRequest Request Object
type SearchHisMeetingsRequest struct {

	// 用户的UUID。 > 该参数将废弃，请勿使用。
	UserUUID *string `json:"userUUID,omitempty"`

	// 查询偏移量。默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。默认是20，最大500条。
	Limit *int32 `json:"limit,omitempty"`

	// 查询条件。会议主题、会议预约人和会议ID等可作为搜索内容。
	SearchKey *string `json:"searchKey,omitempty"`

	// 是否查询企业下所有用户的历史会议。 * true：查询所有用户的历史会议 * false：仅查询管理员自己的历史会议 > 仅对企业管理员生效。
	QueryAll *bool `json:"queryAll,omitempty"`

	// 查询的起始时间戳（单位毫秒）。
	StartDate int64 `json:"startDate"`

	// 查询的截止时间戳（单位毫秒）。
	EndDate int64 `json:"endDate"`

	// 查询结果排序类型。 * ASC_StartTIME：根据会议开始时间升序排序 * DSC_StartTIME：根据会议开始时间降序排序 * ASC_RecordTYPE：根据是否具有录播文件排序，之后默认按照会议开始时间升序排序 * DSC_RecordTYPE：根据是否含有录播文件排序，之后默认按照会议开始时间降序排序
	SortType *string `json:"sortType,omitempty"`

	// 标识是否为第三方portal过来的请求。 > 该参数将废弃，请勿使用。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty"`

	// 用于区分到哪个HCSO站点鉴权。 > 该参数将废弃，请勿使用。
	XSiteId *string `json:"X-Site-Id,omitempty"`
}

func (o SearchHisMeetingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchHisMeetingsRequest struct{}"
	}

	return strings.Join([]string{"SearchHisMeetingsRequest", string(data)}, " ")
}
