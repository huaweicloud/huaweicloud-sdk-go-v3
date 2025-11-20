package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncReplicationMetadataResponse Response Object
type SyncReplicationMetadataResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SyncReplicationMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncReplicationMetadataResponse struct{}"
	}

	return strings.Join([]string{"SyncReplicationMetadataResponse", string(data)}, " ")
}
