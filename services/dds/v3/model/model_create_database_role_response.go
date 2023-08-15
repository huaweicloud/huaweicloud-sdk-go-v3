package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseRoleResponse Response Object
type CreateDatabaseRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRoleResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRoleResponse", string(data)}, " ")
}
