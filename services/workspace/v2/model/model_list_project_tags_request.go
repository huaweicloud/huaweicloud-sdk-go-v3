package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTagsRequest Request Object
type ListProjectTagsRequest struct {

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，取值范围0-1000，默认值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询指定标签key查询，key与value同时传入时，满足其中一个也能查询。
	Key *string `json:"key,omitempty"`

	// 查询指定标签value查询，key与value同时传入时，满足其中一个也能查询。
	Value *string `json:"value,omitempty"`
}

func (o ListProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}
