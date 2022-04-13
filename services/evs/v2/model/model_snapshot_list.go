package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 快照列表信息。
type SnapshotList struct {
	// 云硬盘快照ID。

	Id string `json:"id"`
	// 云硬盘快照的状态。

	Status string `json:"status"`
	// 云硬盘快照名称。

	Name *string `json:"name,omitempty"`
	// 云硬盘快照描述信息。

	Description *string `json:"description,omitempty"`
	// 云硬盘快照创建时间。

	CreatedAt string `json:"created_at"`
	// 云硬盘快照更新时间。

	UpdatedAt *string `json:"updated_at,omitempty"`
	// 云硬盘快照的元数据信息。

	Metadata map[string]string `json:"metadata,omitempty"`
	// 快照所属的云硬盘。

	VolumeId string `json:"volume_id"`
	// 云硬盘快照大小。

	Size int32 `json:"size"`
	// 项目ID。

	OsExtendedSnapshotAttributesprojectId string `json:"os-extended-snapshot-attributes:project_id"`
	// 快照进度。

	OsExtendedSnapshotAttributesprogress string `json:"os-extended-snapshot-attributes:progress"`
	// 专属存储ID。

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`
	// 专属存储名称。

	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`
	// 服务类型。

	ServiceType string `json:"service_type"`
}

func (o SnapshotList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotList struct{}"
	}

	return strings.Join([]string{"SnapshotList", string(data)}, " ")
}
