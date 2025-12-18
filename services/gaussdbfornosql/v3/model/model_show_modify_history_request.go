package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModifyHistoryRequest Request Object
type ShowModifyHistoryRequest struct {

	// **参数解释：** 实例ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 参数名称，支持模糊查询。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ParameterName *string `json:"parameter_name,omitempty"`

	// **参数解释：** 索引位置，偏移量。 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。 **约束限制：** 取值必须为数字，不能为负数。 **取值范围：** 非负整数。 **默认取值：** 0
	Offset *string `json:"offset,omitempty"`

	// **参数解释：** 查询个数上限值。 **约束限制：** 不涉及。 **取值范围：** 1~100。 **默认取值：** 100。不传该参数时，默认查询前100条信息。
	Limit *string `json:"limit,omitempty"`
}

func (o ShowModifyHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModifyHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowModifyHistoryRequest", string(data)}, " ")
}
