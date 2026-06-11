package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterRecoverySettingsRequest Request Object
type ShowDisasterRecoverySettingsRequest struct {

	// **参数解释：** 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：** 索引位置，偏移量。 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。 **约束限制：** 取值必须为数字，不能为负数。 **取值范围：** 非负整数。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 查询个数上限值。 **约束限制：** 不涉及。 **取值范围：** 1~50。 **默认取值：** 50。不传该参数时，默认查询前50条信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDisasterRecoverySettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterRecoverySettingsRequest struct{}"
	}

	return strings.Join([]string{"ShowDisasterRecoverySettingsRequest", string(data)}, " ")
}
