package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAutopilotClusterMasterSnapshotResponse Response Object
type CreateAutopilotClusterMasterSnapshotResponse struct {

	// **参数解释：** 任务ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Uid *string `json:"uid,omitempty"`

	Metadata       *SnapshotCluserResponseMetadata `json:"metadata,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CreateAutopilotClusterMasterSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutopilotClusterMasterSnapshotResponse struct{}"
	}

	return strings.Join([]string{"CreateAutopilotClusterMasterSnapshotResponse", string(data)}, " ")
}
