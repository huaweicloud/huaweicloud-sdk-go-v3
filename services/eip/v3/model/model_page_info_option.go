package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页页码信息
type PageInfoOption struct {

	// 翻页时，作为前一页的marker取值
	PreviousMarker *string `json:"previous_marker,omitempty" xml:"previous_marker"`

	// 翻页时，作为后一页的marker取值
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker"`

	// 当前页的数据总数
	CurrentCount *int32 `json:"current_count,omitempty" xml:"current_count"`
}

func (o PageInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoOption struct{}"
	}

	return strings.Join([]string{"PageInfoOption", string(data)}, " ")
}
