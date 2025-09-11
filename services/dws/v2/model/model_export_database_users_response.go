package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDatabaseUsersResponse Response Object
type ExportDatabaseUsersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportDatabaseUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDatabaseUsersResponse struct{}"
	}

	return strings.Join([]string{"ExportDatabaseUsersResponse", string(data)}, " ")
}
