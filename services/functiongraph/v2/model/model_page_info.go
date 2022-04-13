package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页结构体。
type PageInfo struct {
	// 返回下一页查询地址。

	NextMarker int64 `json:"next_marker"`
	// 返回前一页查询地址。

	PreviousMarker int64 `json:"previous_marker"`
	// 本页返回条目数量。

	CurrentCount int64 `json:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
