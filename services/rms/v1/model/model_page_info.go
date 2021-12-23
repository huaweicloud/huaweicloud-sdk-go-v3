package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页对象
type PageInfo struct {
	// 当前页的资源数量

	CurrentCount *int32 `json:"current_count,omitempty"`
	// 下一页的marker

	NextMarker *string `json:"next_marker,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
