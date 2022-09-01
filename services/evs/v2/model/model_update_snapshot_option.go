package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSnapshotOption struct {

	// 云硬盘快照描述。最大支持255个字节。
	Description *string `json:"description,omitempty" xml:"description"`

	// 云硬盘快照名称。最大支持255个字节。
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o UpdateSnapshotOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotOption struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotOption", string(data)}, " ")
}
