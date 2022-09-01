package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateOnlineMigrationTaskRequest struct {
	Body *CreateOnlineMigrationTaskBody `json:"body,omitempty" xml:"body"`
}

func (o CreateOnlineMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOnlineMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateOnlineMigrationTaskRequest", string(data)}, " ")
}
