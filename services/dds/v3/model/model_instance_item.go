package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceItem struct {

	// **参数解释：** 实例ID。 **取值范围：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 实例名称 **取值范围：** 不涉及。
	InstanceName string `json:"instance_name"`

	// **参数解释：** 标签列表。如果没有标签，默认为空数组。 **取值范围：** 不涉及。
	Tags []InstanceItemTagItem `json:"tags"`
}

func (o InstanceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceItem struct{}"
	}

	return strings.Join([]string{"InstanceItem", string(data)}, " ")
}
