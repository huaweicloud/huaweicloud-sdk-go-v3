package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMigrationTaskRequest Request Object
type DeleteMigrationTaskRequest struct {
	Body *DeleteMigrateTaskRequest `json:"body,omitempty"`
}

func (o DeleteMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMigrationTaskRequest", string(data)}, " ")
}
