package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSnapshotGroupTagRequest Request Object
type BatchCreateSnapshotGroupTagRequest struct {

	// **参数解释** 快照组ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	SnapshotGroupId string `json:"snapshot_group_id"`

	Body *BatchCreateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSnapshotGroupTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSnapshotGroupTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSnapshotGroupTagRequest", string(data)}, " ")
}
