package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotRegion 快照区域信息
type SnapshotRegion struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`
}

func (o SnapshotRegion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotRegion struct{}"
	}

	return strings.Join([]string{"SnapshotRegion", string(data)}, " ")
}
