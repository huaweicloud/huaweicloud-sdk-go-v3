package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSnapshotTagsRequest Request Object
type BatchDeleteSnapshotTagsRequest struct {

	// **参数解释** 快照ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	SnapshotId string `json:"snapshot_id"`

	Body *BatchDeleteResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSnapshotTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSnapshotTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSnapshotTagsRequest", string(data)}, " ")
}
