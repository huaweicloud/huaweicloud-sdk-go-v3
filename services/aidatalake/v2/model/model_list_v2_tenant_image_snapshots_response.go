package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2TenantImageSnapshotsResponse Response Object
type ListV2TenantImageSnapshotsResponse struct {

	// **参数解释**：快照总数。 **约束限制**：不涉及。 **取值范围**：0~10000。 **默认取值**：不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：镜像快照列表。
	ImageSnapshots *[]V2ImageSnapshotInfo `json:"image_snapshots,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListV2TenantImageSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2TenantImageSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListV2TenantImageSnapshotsResponse", string(data)}, " ")
}
