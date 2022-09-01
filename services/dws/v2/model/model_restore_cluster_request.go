package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestoreClusterRequest struct {

	// 待恢复的快照ID。
	SnapshotId string `json:"snapshot_id" xml:"snapshot_id"`

	Body *RestoreClusterRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RestoreClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreClusterRequest struct{}"
	}

	return strings.Join([]string{"RestoreClusterRequest", string(data)}, " ")
}
