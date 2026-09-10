package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubNetworkInterfacesByTagsRequest Request Object
type ListSubNetworkInterfacesByTagsRequest struct {

	// **参数解释**： 查询记录数。 **约束限制**： 不涉及。 **取值范围**： 1-1000 **默认取值**： 1000
	Limit *string `json:"limit,omitempty"`

	// **参数解释**： 索引位置，从第一条数据偏移offset条数据后开始查询。 **约束限制**： 必须为数字，不能为负数。 **取值范围**： 不涉及。 **默认取值**： 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	Body *ListSubNetworkInterfacesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListSubNetworkInterfacesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubNetworkInterfacesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSubNetworkInterfacesByTagsRequest", string(data)}, " ")
}
