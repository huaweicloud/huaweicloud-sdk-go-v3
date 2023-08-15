package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchQosHistoryMeetingsRequest Request Object
type SearchQosHistoryMeetingsRequest struct {

	// 查询的起始日期，Unix时间戳（单位毫秒）。
	StartDate int64 `json:"startDate"`

	// 查询的截止日期，Unix时间戳（单位毫秒）。
	EndDate int64 `json:"endDate"`

	// 查询偏移量。 * 取值：大于等于0，默认值为0 * 大于等于最大条目数量，则返回最后一页的数据
	Offset *int32 `json:"offset,omitempty"`

	// 查询的条目数量。 * 取值：1-500，默认值为20
	Limit *int32 `json:"limit,omitempty"`

	// 查询条件。会议主题、会议预约人和会议ID等可作为搜索内容。长度限制为1-512个字符。
	SearchKey *string `json:"searchKey,omitempty"`
}

func (o SearchQosHistoryMeetingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosHistoryMeetingsRequest struct{}"
	}

	return strings.Join([]string{"SearchQosHistoryMeetingsRequest", string(data)}, " ")
}
