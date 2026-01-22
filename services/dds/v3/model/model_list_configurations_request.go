package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsRequest Request Object
type ListConfigurationsRequest struct {

	// **参数解释：** 索引位置，偏移量。 **约束限制：** 必须为数字。 **取值范围：** 不能为负数。 **默认取值：** 0。偏移0条数据，表示从第一条数据开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 每页显示的数量。 **约束限制：** 不涉及。 **取值范围：** 1~100。 **默认取值：** 100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationsRequest", string(data)}, " ")
}
