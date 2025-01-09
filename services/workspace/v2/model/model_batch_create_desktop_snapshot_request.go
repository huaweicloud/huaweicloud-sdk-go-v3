package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDesktopSnapshotRequest Request Object
type BatchCreateDesktopSnapshotRequest struct {
	Body *BatchCreateDesktopSnapshotReq `json:"body,omitempty"`
}

func (o BatchCreateDesktopSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDesktopSnapshotRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDesktopSnapshotRequest", string(data)}, " ")
}
