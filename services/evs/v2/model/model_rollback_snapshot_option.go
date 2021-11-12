package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RollbackSnapshotOption struct {
	// 回滚的目标云硬盘名称。

	Name *string `json:"name,omitempty"`
	// 回滚的目标云硬盘UUID。

	VolumeId string `json:"volume_id"`
}

func (o RollbackSnapshotOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotOption struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotOption", string(data)}, " ")
}
