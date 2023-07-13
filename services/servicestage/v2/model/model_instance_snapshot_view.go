package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceSnapshotView 快照参数。
type InstanceSnapshotView struct {

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 应用组件实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`
}

func (o InstanceSnapshotView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSnapshotView struct{}"
	}

	return strings.Join([]string{"InstanceSnapshotView", string(data)}, " ")
}
