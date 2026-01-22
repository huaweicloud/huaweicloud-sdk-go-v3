package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppliedInstancesRequest Request Object
type ListAppliedInstancesRequest struct {

	// **参数解释：** 参数模板ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ConfigId string `json:"config_id"`

	// **参数解释：** 索引位置，偏移量。 **约束限制：** 必须为整数数字。 **取值范围：** >=0。 **默认取值：** 0。偏移0条数据，表示从第一条数据开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 查询记录数。 **约束限制：** 正整数。 **取值范围：** 1~100。 **默认取值：** 100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAppliedInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppliedInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAppliedInstancesRequest", string(data)}, " ")
}
