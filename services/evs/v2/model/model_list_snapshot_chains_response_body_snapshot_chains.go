package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSnapshotChainsResponseBodySnapshotChains struct {

	// **参数解释** 快照链ID。 **取值范围** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释** 快照链所属云硬盘的可用区。 **取值范围** 不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释** 快照链上的快照个数。 **取值范围** 大于等于0。
	SnapshotCount *int32 `json:"snapshot_count,omitempty"`

	// **参数解释** 快照链的大小。 **取值范围** 大于等于0。
	Capacity *float32 `json:"capacity,omitempty"`

	// **参数解释** 项目ID。 **取值范围** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释** 快照链所属的云硬盘ID。 **取值范围** 不涉及。
	VolumeId *string `json:"volume_id,omitempty"`

	// **参数解释** 快照链的类型。 **取值范围** 当前快照链支持\"standard\"、\"standard_copy\"两种类型。
	Category *string `json:"category,omitempty"`

	// **参数解释** 创建时间。 **取值范围** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释** 更新时间。 **取值范围** 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ListSnapshotChainsResponseBodySnapshotChains) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotChainsResponseBodySnapshotChains struct{}"
	}

	return strings.Join([]string{"ListSnapshotChainsResponseBodySnapshotChains", string(data)}, " ")
}
