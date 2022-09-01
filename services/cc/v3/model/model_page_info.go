package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页查询页的信息。
type PageInfo struct {

	// 下一页的marker，值为资源的uuid，为空时表示最后一页。
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker"`

	// 上一页的marker，值为资源的uuid，为空时表示第一页。
	PreviousMarker *string `json:"previous_marker,omitempty" xml:"previous_marker"`

	// 当前列表中资源数量。
	CurrentCount *int32 `json:"current_count,omitempty" xml:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
