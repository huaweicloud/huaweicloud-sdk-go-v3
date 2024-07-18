package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageInfo 分页信息
type PageInfo struct {

	// 参数解释：分页查询结果中第一条记录的ID。通常用于配合page_reverse=true查询上一页。
	PreviousMarker string `json:"previous_marker"`

	// 参数解释：分页查询结果中最后一条记录的ID。通常用于查询下一页。
	NextMarker *string `json:"next_marker,omitempty"`

	// 参数解释：当前的记录数。
	CurrentCount int32 `json:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
