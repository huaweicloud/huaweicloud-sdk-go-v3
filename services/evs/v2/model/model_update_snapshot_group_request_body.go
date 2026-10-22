package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSnapshotGroupRequestBody struct {
	SnapshotGroup *UpdateSnapshotGroupResponseBody `json:"snapshot_group"`
}

func (o UpdateSnapshotGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotGroupRequestBody", string(data)}, " ")
}
