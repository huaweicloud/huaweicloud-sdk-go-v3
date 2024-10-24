package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFabricProjectTagsRequest Request Object
type ListFabricProjectTagsRequest struct {

	// 使用预定义标签，true表示使用。
	UsePredefineTags bool `json:"use_predefine_tags"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数默认为1000，limit最多为1000,不能为负数，最小值为1
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFabricProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFabricProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListFabricProjectTagsRequest", string(data)}, " ")
}
