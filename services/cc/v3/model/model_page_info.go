package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageInfo 分页查询页的信息。
type PageInfo struct {

	// 向后分页标识符。
	NextMarker *string `json:"next_marker,omitempty"`

	// 向前分页标识符。
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// 当前列表中资源数量。
	CurrentCount int32 `json:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
