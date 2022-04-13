package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSnapshotRequest struct {
	Body *CreateSnapshotRequestBody `json:"body,omitempty"`
}

func (o CreateSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotRequest struct{}"
	}

	return strings.Join([]string{"CreateSnapshotRequest", string(data)}, " ")
}
