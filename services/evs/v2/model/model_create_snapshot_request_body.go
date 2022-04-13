package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateSnapshotRequestBody struct {
	Snapshot *CreateSnapshotOption `json:"snapshot"`
}

func (o CreateSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSnapshotRequestBody", string(data)}, " ")
}
