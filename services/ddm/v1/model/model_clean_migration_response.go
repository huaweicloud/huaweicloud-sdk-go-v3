package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CleanMigrationResponse Response Object
type CleanMigrationResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CleanMigrationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CleanMigrationResponse struct{}"
	}

	return strings.Join([]string{"CleanMigrationResponse", string(data)}, " ")
}
