package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 恢复集群
type RestorePoint struct {

	// 快照ID
	BackRef *string `json:"back_ref,omitempty"`

	// 恢复时间
	RestoreTime *int64 `json:"restore_time,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o RestorePoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePoint struct{}"
	}

	return strings.Join([]string{"RestorePoint", string(data)}, " ")
}
