package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
