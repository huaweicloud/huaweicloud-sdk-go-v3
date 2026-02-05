package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNonMigrationUsersResponse Response Object
type UpdateNonMigrationUsersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNonMigrationUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNonMigrationUsersResponse struct{}"
	}

	return strings.Join([]string{"UpdateNonMigrationUsersResponse", string(data)}, " ")
}
