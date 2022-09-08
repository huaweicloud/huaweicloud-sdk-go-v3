package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchQosOnlineMeetingsRequest struct {

	// 查询偏移量。 * 取值：大于等于0，默认值为0。 * 小于最小值0时，系统设置为0。 * 大于等于最大条目数量，则返回最后一页的数据。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的条目数量。 * 取值：1-500，默认值为20。 * 小于最小值1时，系统设置为1。 * 大于最大值500时，系统设置为500。
	Limit *int32 `json:"limit,omitempty"`

	// 根据会议主题,预定人和会议id作为关键词，模糊查询会议列表。最大不超过512个字节。
	SearchKey *string `json:"searchKey,omitempty"`
}

func (o SearchQosOnlineMeetingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosOnlineMeetingsRequest struct{}"
	}

	return strings.Join([]string{"SearchQosOnlineMeetingsRequest", string(data)}, " ")
}
