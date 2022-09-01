package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSnapshotRequest struct {
	Body *CreateSnapshotRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotRequest struct{}"
	}

	return strings.Join([]string{"CreateSnapshotRequest", string(data)}, " ")
}
