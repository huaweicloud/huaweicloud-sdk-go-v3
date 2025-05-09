package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageInfo 分页信息。
type PageInfo struct {

	// 下一页的页面标识。
	NextMarker *string `json:"next_marker,omitempty"`

	// 上一页的页面标识。
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// 页面数量。
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
