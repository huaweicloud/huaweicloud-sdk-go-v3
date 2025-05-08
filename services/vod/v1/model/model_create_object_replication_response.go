package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectReplicationResponse Response Object
type CreateObjectReplicationResponse struct {

	// 任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateObjectReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectReplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateObjectReplicationResponse", string(data)}, " ")
}
