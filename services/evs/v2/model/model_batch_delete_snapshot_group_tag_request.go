package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSnapshotGroupTagRequest Request Object
type BatchDeleteSnapshotGroupTagRequest struct {

	// **参数解释** 快照组ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	SnapshotGroupId string `json:"snapshot_group_id"`

	Body *BatchDeleteResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSnapshotGroupTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSnapshotGroupTagRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSnapshotGroupTagRequest", string(data)}, " ")
}
