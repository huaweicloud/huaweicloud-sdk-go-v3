package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreClusterRequest Request Object
type RestoreClusterRequest struct {

	// 待恢复的快照ID。
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
