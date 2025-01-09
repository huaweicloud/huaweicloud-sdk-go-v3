package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopSnapshotResponse Response Object
type BatchDeleteDesktopSnapshotResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteDesktopSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopSnapshotResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopSnapshotResponse", string(data)}, " ")
}
