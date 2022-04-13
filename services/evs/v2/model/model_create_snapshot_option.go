package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 快照信息。
type CreateSnapshotOption struct {
	// 源云硬盘的ID。

	VolumeId string `json:"volume_id"`
	// 强制创快照标示，默认为false。 当force标记为false时，云硬盘处于挂载状态时，不能强制创建快照。 当force标记为true时，即使云硬盘处于挂载状态时，仍可以创建快照。

	Force *bool `json:"force,omitempty"`
	// 云硬盘快照的元数据信息。

	Metadata map[string]string `json:"metadata,omitempty"`
	// 云硬盘快照描述，最大支持255个字节。

	Description *string `json:"description,omitempty"`
	// 云硬盘快照名称。最大支持255个字节。  > > 说明： > 对云硬盘创建备份时，同时会创建以autobk_snapshot_为名称前缀的快照，云硬盘控制台对此类快照会有操作限制。因此建议不要创建以> > autobk_snapshot_为名称前缀的快照，避免影响快照的正常使用

	Name *string `json:"name,omitempty"`
}

func (o CreateSnapshotOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotOption struct{}"
	}

	return strings.Join([]string{"CreateSnapshotOption", string(data)}, " ")
}
