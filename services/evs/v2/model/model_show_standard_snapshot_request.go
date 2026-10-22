package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStandardSnapshotRequest Request Object
type ShowStandardSnapshotRequest struct {

	// **参数解释** 快照ID。 可通过[查询云硬盘快照详情列表](evs_04_2021.xml)接口获取云硬盘快照ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	SnapshotId string `json:"snapshot_id"`
}

func (o ShowStandardSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStandardSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ShowStandardSnapshotRequest", string(data)}, " ")
}
