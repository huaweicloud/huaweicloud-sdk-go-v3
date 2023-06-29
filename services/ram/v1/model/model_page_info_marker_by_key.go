package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageInfoMarkerByKey 分页详细信息。
type PageInfoMarkerByKey struct {

	// 上一页的页面标识。
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// 下一页的页面标识。
	NextMarker *string `json:"next_marker,omitempty"`

	// 页面数量。
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o PageInfoMarkerByKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoMarkerByKey struct{}"
	}

	return strings.Join([]string{"PageInfoMarkerByKey", string(data)}, " ")
}
