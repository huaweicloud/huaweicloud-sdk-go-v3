package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSnapshotTagsRequest Request Object
type BatchCreateSnapshotTagsRequest struct {

	// **参数解释** 快照ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	SnapshotId string `json:"snapshot_id"`

	Body *BatchCreateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSnapshotTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSnapshotTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSnapshotTagsRequest", string(data)}, " ")
}
