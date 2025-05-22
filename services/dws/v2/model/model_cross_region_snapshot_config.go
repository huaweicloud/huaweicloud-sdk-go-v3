package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CrossRegionSnapshotConfig 快照跨区域配置信息
type CrossRegionSnapshotConfig struct {

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 源区域
	SourceRegion *string `json:"source_region,omitempty"`

	// 源项目ID
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// 目的区域
	DestinationRegion *string `json:"destination_region,omitempty"`

	// 目的项目ID
	DestinationProjectId *string `json:"destination_project_id,omitempty"`

	// 状态
	Status *bool `json:"status,omitempty"`

	// 保存时间
	BackKeepDay *int32 `json:"back_keep_day,omitempty"`

	// 总大小
	TotalSize *int64 `json:"total_size,omitempty"`
}

func (o CrossRegionSnapshotConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CrossRegionSnapshotConfig struct{}"
	}

	return strings.Join([]string{"CrossRegionSnapshotConfig", string(data)}, " ")
}
