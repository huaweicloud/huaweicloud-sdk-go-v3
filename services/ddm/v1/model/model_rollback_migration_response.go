package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackMigrationResponse Response Object
type RollbackMigrationResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RollbackMigrationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackMigrationResponse struct{}"
	}

	return strings.Join([]string{"RollbackMigrationResponse", string(data)}, " ")
}
