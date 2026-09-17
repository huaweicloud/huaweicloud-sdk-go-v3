package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterMasterSnapshotResponse Response Object
type CreateClusterMasterSnapshotResponse struct {

	// **参数解释：** 任务ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Uid *string `json:"uid,omitempty"`

	Metadata       *SnapshotCluserResponseMetadata `json:"metadata,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CreateClusterMasterSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterMasterSnapshotResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterMasterSnapshotResponse", string(data)}, " ")
}
