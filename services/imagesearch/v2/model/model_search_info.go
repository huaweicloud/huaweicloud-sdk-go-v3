package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchInfo 搜索结果的相关信息。
type SearchInfo struct {

	// 搜索结果总数。
	TotalNum float32 `json:"total_num,omitempty"`

	// 返回结果总数。
	ReturnNum float32 `json:"return_num,omitempty"`

	// 搜索过程耗时，单位为毫秒。
	SearchTime float32 `json:"search_time,omitempty"`

	LastItem *SearchAfterParam `json:"last_item,omitempty"`
}

func (o SearchInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchInfo struct{}"
	}

	return strings.Join([]string{"SearchInfo", string(data)}, " ")
}
