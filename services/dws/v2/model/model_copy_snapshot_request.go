package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopySnapshotRequest Request Object
type CopySnapshotRequest struct {

	// **参数解释**： 快照ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SnapshotId string `json:"snapshot_id"`

	Body *LinkCopyReq `json:"body,omitempty"`
}

func (o CopySnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopySnapshotRequest struct{}"
	}

	return strings.Join([]string{"CopySnapshotRequest", string(data)}, " ")
}
