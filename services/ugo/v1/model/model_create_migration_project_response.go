package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateMigrationProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMigrationProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMigrationProjectResponse struct{}"
	}

	return strings.Join([]string{"CreateMigrationProjectResponse", string(data)}, " ")
}
