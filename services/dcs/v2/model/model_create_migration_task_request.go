package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMigrationTaskRequest struct {
	Body *CreateMigrationTaskBody `json:"body,omitempty"`
}

func (o CreateMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateMigrationTaskRequest", string(data)}, " ")
}
