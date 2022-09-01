package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateSnapshotRequestBody struct {
	Snapshot *UpdateSnapshotOption `json:"snapshot" xml:"snapshot"`
}

func (o UpdateSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotRequestBody", string(data)}, " ")
}
