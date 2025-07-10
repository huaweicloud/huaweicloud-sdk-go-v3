package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopByTagsResponse Response Object
type ListDesktopByTagsResponse struct {

	// 指定查询信息列表的偏移量，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 资源对象。
	Resources *[]TagResource `json:"resources,omitempty"`

	// 数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopByTagsResponse", string(data)}, " ")
}
