package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 翻页信息
type PageInfo struct {
	// 分页查询结果中第一条记录的ID

	PreviousMarker string `json:"previous_marker"`
	// 分页查询结果中最后一条记录的ID。

	NextMarker *string `json:"next_marker,omitempty"`
	// 当前的记录数。

	CurrentCount int32 `json:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
