package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageInfo 分页信息。
type PageInfo struct {

	// 翻页页数，从1开始。
	Offset *int32 `json:"offset,omitempty"`

	// 每页展示的条数。
	Limit *int32 `json:"limit,omitempty"`

	// 总条数。
	Total *int32 `json:"total,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
