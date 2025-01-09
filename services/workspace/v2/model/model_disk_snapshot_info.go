package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiskSnapshotInfo 磁盘快照信息。
type DiskSnapshotInfo struct {

	// 磁盘快照名称，批量操作时最终名称为：桌面名称+磁盘快照名称+时间戳。
	Name *string `json:"name,omitempty"`

	// 磁盘快照描述。
	Description *string `json:"description,omitempty"`
}

func (o DiskSnapshotInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskSnapshotInfo struct{}"
	}

	return strings.Join([]string{"DiskSnapshotInfo", string(data)}, " ")
}
