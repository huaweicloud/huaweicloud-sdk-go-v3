package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionOnComponentSpec struct {

	// 快照序号。
	SnapshotIndex *int32 `json:"snapshot_index,omitempty"`

	// 实例个数。
	Replica *int32 `json:"replica,omitempty"`

	Source *ActionOnComponentSource `json:"source,omitempty"`

	ResourceLimit *ResourceLimitForUpgrade `json:"resource_limit,omitempty"`
}

func (o ActionOnComponentSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionOnComponentSpec struct{}"
	}

	return strings.Join([]string{"ActionOnComponentSpec", string(data)}, " ")
}
