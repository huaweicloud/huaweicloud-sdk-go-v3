package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopSnapshotRequest Request Object
type BatchDeleteDesktopSnapshotRequest struct {
	Body *BatchDeleteDesktopSnapshotReq `json:"body,omitempty"`
}

func (o BatchDeleteDesktopSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopSnapshotRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopSnapshotRequest", string(data)}, " ")
}
