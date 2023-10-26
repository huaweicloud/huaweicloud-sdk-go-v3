package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMigrationTaskResponse Response Object
type UpdateMigrationTaskResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateMigrationTaskResponse", string(data)}, " ")
}
