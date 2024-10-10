package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReplicationJobResponse Response Object
type CreateReplicationJobResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReplicationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReplicationJobResponse struct{}"
	}

	return strings.Join([]string{"CreateReplicationJobResponse", string(data)}, " ")
}
