package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreLogReplayDatabaseResponse Response Object
type RestoreLogReplayDatabaseResponse struct {

	// 任务流id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreLogReplayDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreLogReplayDatabaseResponse struct{}"
	}

	return strings.Join([]string{"RestoreLogReplayDatabaseResponse", string(data)}, " ")
}
