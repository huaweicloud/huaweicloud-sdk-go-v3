package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CrossRegionSnapshotConfig **参数解释**： 快照跨区域配置信息。 **取值范围**： 不涉及。
type CrossRegionSnapshotConfig struct {

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群名称。 **取值范围**： 不涉及。
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 源区域。 **取值范围**： 不涉及。
	SourceRegion *string `json:"source_region,omitempty"`

	// **参数解释**： 源项目ID。 **取值范围**： 不涉及。
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// **参数解释**： 目的区域。 **取值范围**： 不涉及。
	DestinationRegion *string `json:"destination_region,omitempty"`

	// **参数解释**： 目的项目ID。 **取值范围**： 不涉及。
	DestinationProjectId *string `json:"destination_project_id,omitempty"`

	// **参数解释**： 状态。 **取值范围**： 不涉及。
	Status *bool `json:"status,omitempty"`

	// **参数解释**： 保存时间。 **取值范围**： 不涉及。
	BackKeepDay *int32 `json:"back_keep_day,omitempty"`

	// **参数解释**： 总大小。 **取值范围**： 大于等于0。
	TotalSize *int64 `json:"total_size,omitempty"`
}

func (o CrossRegionSnapshotConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CrossRegionSnapshotConfig struct{}"
	}

	return strings.Join([]string{"CrossRegionSnapshotConfig", string(data)}, " ")
}
