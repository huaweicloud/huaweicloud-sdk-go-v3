package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSnapshotRequest Request Object
type CheckSnapshotRequest struct {
	Body *CheckSnapshotReq `json:"body,omitempty"`
}

func (o CheckSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSnapshotRequest struct{}"
	}

	return strings.Join([]string{"CheckSnapshotRequest", string(data)}, " ")
}
