package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMigrationProjectRequest Request Object
type CreateMigrationProjectRequest struct {
	Body *CreateMigrationProject `json:"body,omitempty"`
}

func (o CreateMigrationProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMigrationProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateMigrationProjectRequest", string(data)}, " ")
}
