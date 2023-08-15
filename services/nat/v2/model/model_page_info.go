package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageInfo 分页信息。
type PageInfo struct {

	// 分页查询结果中最后一条记录的ID。通常用于查询下一页。
	NextMarker *string `json:"next_marker,omitempty"`

	// 分页查询结果中第一条记录的ID。通常用于配合page_reverse=true查询上一页。
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// 分页查询资源时，本页的实例的个数。
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
