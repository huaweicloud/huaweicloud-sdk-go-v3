package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsByTagsResponse Response Object
type ListDesktopsByTagsResponse struct {

	// 指定查询信息列表的偏移量，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 资源对象。
	Resources *[]TagResource `json:"resources,omitempty"`

	// 数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopsByTagsResponse", string(data)}, " ")
}
