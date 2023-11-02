package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTableRestoreRequest Request Object
type CheckTableRestoreRequest struct {

	// 快照ID
	SnapshotId string `json:"snapshot_id"`

	Body *CheckTableRestoreRequestBody `json:"body,omitempty"`
}

func (o CheckTableRestoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTableRestoreRequest struct{}"
	}

	return strings.Join([]string{"CheckTableRestoreRequest", string(data)}, " ")
}
