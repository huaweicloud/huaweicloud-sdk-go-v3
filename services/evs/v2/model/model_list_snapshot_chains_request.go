package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotChainsRequest Request Object
type ListSnapshotChainsRequest struct {

	// **参数解释** 分页查询的起始资源ID，取值为上一页最后一条查询记录的资源ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释** 返回结果个数限制。 **约束限制** 不涉及。 **取值范围** 最小值1，最大值1000，默认为1000。返回的结果中记录数不超过limit值。 **默认取值** 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释** 快照链ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释** 快照链所属的云硬盘ID。 可通过[查询所有云硬盘详情](evs_04_2006.xml)获取云硬盘ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	VolumeId *string `json:"volume_id,omitempty"`

	// **参数解释** 快照链的类型。 **约束限制** 不涉及。 **取值范围** 当前快照链支持\"standard\"、\"backup\"和\"server_backup\"三种类型。 **默认取值** 不涉及。
	Category *string `json:"category,omitempty"`
}

func (o ListSnapshotChainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotChainsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotChainsRequest", string(data)}, " ")
}
