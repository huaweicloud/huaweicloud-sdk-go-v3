package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSnapshotChainsRequest Request Object
type CountSnapshotChainsRequest struct {

	// **参数解释** 快照链所属的云硬盘ID。 可通过[查询所有云硬盘详情](evs_04_2006.xml)获取云硬盘ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	VolumeId *string `json:"volume_id,omitempty"`

	// **参数解释** 快照链对应的云硬盘所在的可用区。  **约束限制** 不涉及。 **取值范围** 不涉及。         **默认取值** 不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释** 快照链的类型。 **约束限制** 不涉及。 **取值范围** 当前快照链支持\"standard\"、\"standard_copy\"两种类型。  **默认取值** 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释** 快照链ID。 可通过[查询快照链列表](ListSnapshotChains.xml)接口获取快照链ID。 **约束限制** 不涉及。 **取值范围** 不涉及。  **默认取值** 不涉及。
	Id *string `json:"id,omitempty"`
}

func (o CountSnapshotChainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSnapshotChainsRequest struct{}"
	}

	return strings.Join([]string{"CountSnapshotChainsRequest", string(data)}, " ")
}
