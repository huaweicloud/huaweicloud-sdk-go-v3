package model

import (
	"encoding/json"

	"strings"
)

type UpdateSnapshotOption struct {
	// 云硬盘快照描述。最大支持255个字节。

	Description *string `json:"description,omitempty"`
	// 云硬盘快照名称。最大支持255个字节。

	Name *string `json:"name,omitempty"`
}

func (o UpdateSnapshotOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSnapshotOption struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotOption", string(data)}, " ")
}
