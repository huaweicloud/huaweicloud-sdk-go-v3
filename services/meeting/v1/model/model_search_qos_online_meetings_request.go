package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchQosOnlineMeetingsRequest Request Object
type SearchQosOnlineMeetingsRequest struct {

	// 查询偏移量。 * 取值：大于等于0，默认值为0 * 大于等于最大条目数量，则返回最后一页的数据
	Offset *int32 `json:"offset,omitempty"`

	// 查询的条目数量。 * 取值：1-500，默认值为20
	Limit *int32 `json:"limit,omitempty"`

	// 查询条件。会议主题、会议预约人和会议ID等可作为搜索内容。长度限制为1-512个字符。
	SearchKey *string `json:"searchKey,omitempty"`
}

func (o SearchQosOnlineMeetingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosOnlineMeetingsRequest struct{}"
	}

	return strings.Join([]string{"SearchQosOnlineMeetingsRequest", string(data)}, " ")
}
