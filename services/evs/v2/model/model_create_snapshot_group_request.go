package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSnapshotGroupRequest Request Object
type CreateSnapshotGroupRequest struct {
	Body *CreateSnapshotGroupRequestBody `json:"body,omitempty"`
}

func (o CreateSnapshotGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateSnapshotGroupRequest", string(data)}, " ")
}
