package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页信息
type PageInfo struct {

	// 下一页标识。
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker"`

	// 当前页数量。
	CurrentCount *int32 `json:"current_count,omitempty" xml:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
