package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDesktopSnapshotResponse Response Object
type BatchCreateDesktopSnapshotResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateDesktopSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDesktopSnapshotResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDesktopSnapshotResponse", string(data)}, " ")
}
