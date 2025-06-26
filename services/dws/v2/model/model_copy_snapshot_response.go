package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopySnapshotResponse Response Object
type CopySnapshotResponse struct {

	// **参数解释**： 快照ID。 **取值范围**： 不涉及。
	SnapshotId     *string `json:"snapshot_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopySnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopySnapshotResponse struct{}"
	}

	return strings.Join([]string{"CopySnapshotResponse", string(data)}, " ")
}
