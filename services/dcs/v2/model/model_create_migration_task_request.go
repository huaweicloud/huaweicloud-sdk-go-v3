package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMigrationTaskRequest Request Object
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
