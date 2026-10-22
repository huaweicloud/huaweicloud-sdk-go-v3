package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2TenantImageSnapshotsRequest Request Object
type ListV2TenantImageSnapshotsRequest struct {

	// **参数解释**：镜像快照状态。 **约束限制**：不涉及。 **取值范围**：   - Creating：创建中。   - Available：可用。   - Unavailable：不可用。 **默认取值**：不涉及。
	Status *ImageSnapshotStatus `json:"status,omitempty"`

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：分页大小。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListV2TenantImageSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2TenantImageSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"ListV2TenantImageSnapshotsRequest", string(data)}, " ")
}
