package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DualActiveInfo struct {

	// **参数解释：** 双活角色。 **取值范围：** 不涉及。
	Role *string `json:"role,omitempty"`

	// **参数解释：** 双活状态。 **取值范围：** - normal：表示双活状态正常。 - abnormal：表示双活状态异常。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 双活对端实例id。 **取值范围：** 不涉及。
	DestinationInstanceId *string `json:"destination_instance_id,omitempty"`

	// **参数解释：** 双活对端region。 **取值范围：** 不涉及。
	DestinationRegion *string `json:"destination_region,omitempty"`

	// **参数解释：** 双活对端实例名称。 **取值范围：** 不涉及。
	DestinationInstanceName *string `json:"destination_instance_name,omitempty"`

	// **参数解释：** 双活对端实例节点数量。 **取值范围：** 不涉及。
	DestinationInstanceNodeNum *string `json:"destination_instance_node_num,omitempty"`

	// **参数解释：** 双活对端实例规格。 **取值范围：** 不涉及。
	DestinationInstanceSpecCode *string `json:"destination_instance_spec_code,omitempty"`
}

func (o DualActiveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DualActiveInfo struct{}"
	}

	return strings.Join([]string{"DualActiveInfo", string(data)}, " ")
}
