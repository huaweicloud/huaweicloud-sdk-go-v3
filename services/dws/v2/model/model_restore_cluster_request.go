package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreClusterRequest Request Object
type RestoreClusterRequest struct {

	// **参数解释**： 待恢复的快照ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SnapshotId string `json:"snapshot_id"`

	Body *RestoreClusterRequestBody `json:"body,omitempty"`
}

func (o RestoreClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreClusterRequest struct{}"
	}

	return strings.Join([]string{"RestoreClusterRequest", string(data)}, " ")
}
