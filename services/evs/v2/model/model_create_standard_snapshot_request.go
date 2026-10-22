package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStandardSnapshotRequest Request Object
type CreateStandardSnapshotRequest struct {
	Body *CreateStandardSnapshotRequestBody `json:"body,omitempty"`
}

func (o CreateStandardSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStandardSnapshotRequest struct{}"
	}

	return strings.Join([]string{"CreateStandardSnapshotRequest", string(data)}, " ")
}
