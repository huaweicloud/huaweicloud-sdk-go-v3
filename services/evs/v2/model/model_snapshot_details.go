package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 快照详情。
type SnapshotDetails struct {

	// 云硬盘快照ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 云硬盘快照状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 云硬盘快照名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 云硬盘快照描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 云硬盘快照创建时间。 时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 快照更新时间。 时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 云硬盘快照的元数据信息。
	Metadata *interface{} `json:"metadata,omitempty" xml:"metadata"`

	// 快照所属的云硬盘ID。
	VolumeId *string `json:"volume_id,omitempty" xml:"volume_id"`

	// 云硬盘快照大小，单位为GB。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 预留属性。
	OsExtendedSnapshotAttributesprojectId *string `json:"os-extended-snapshot-attributes:project_id,omitempty" xml:"os-extended-snapshot-attributes:project_id"`

	// 预留属性。
	OsExtendedSnapshotAttributesprogress *string `json:"os-extended-snapshot-attributes:progress,omitempty" xml:"os-extended-snapshot-attributes:progress"`
}

func (o SnapshotDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotDetails struct{}"
	}

	return strings.Join([]string{"SnapshotDetails", string(data)}, " ")
}
