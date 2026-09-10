package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBindMigrationResourceToWorkspaceResponse Response Object
type BatchBindMigrationResourceToWorkspaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchBindMigrationResourceToWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindMigrationResourceToWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"BatchBindMigrationResourceToWorkspaceResponse", string(data)}, " ")
}
