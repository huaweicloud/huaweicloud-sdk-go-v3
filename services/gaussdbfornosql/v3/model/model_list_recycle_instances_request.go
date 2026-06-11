package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleInstancesRequest Request Object
type ListRecycleInstancesRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释：** 索引位置，偏移量。 索引位置偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制：** 大于或等于0。 **取值范围：** 不涉及。 **默认取值：** 0，表示从第一条数据开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 查询个数上限值。 **约束限制：** 不涉及。 **取值范围：** 1~100。 **默认取值：** 100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRecycleInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesRequest", string(data)}, " ")
}
