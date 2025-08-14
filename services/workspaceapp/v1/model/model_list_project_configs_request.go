package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectConfigsRequest Request Object
type ListProjectConfigsRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询名称,模糊匹配。
	Name *string `json:"name,omitempty"`
}

func (o ListProjectConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectConfigsRequest", string(data)}, " ")
}
