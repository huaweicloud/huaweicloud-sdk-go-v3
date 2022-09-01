package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群操作进度，任务信息，由key、value组成。key值为正在进行的任务，value值为正在进行任务的进度。示例如 \"action_progress\":{\"SNAPSHOTTING\":\"16%\"}
type ActionProgress struct {

	// 创建集群进度，例如：29%
	Creating *string `json:"CREATING,omitempty" xml:"CREATING"`

	// 扩容集群进度，例如：29%
	Growing *string `json:"GROWING,omitempty" xml:"GROWING"`

	// 恢复集群进度，例如：29%
	Restoring *string `json:"RESTORING,omitempty" xml:"RESTORING"`

	// 集群快照进度，例如：29%
	Snapshotting *string `json:"SNAPSHOTTING,omitempty" xml:"SNAPSHOTTING"`

	// 修复集群进度，例如：29%
	Repairing *string `json:"REPAIRING,omitempty" xml:"REPAIRING"`
}

func (o ActionProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionProgress struct{}"
	}

	return strings.Join([]string{"ActionProgress", string(data)}, " ")
}
