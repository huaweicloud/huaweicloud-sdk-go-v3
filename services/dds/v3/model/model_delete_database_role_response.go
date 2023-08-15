package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseRoleResponse Response Object
type DeleteDatabaseRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRoleResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRoleResponse", string(data)}, " ")
}
