package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbRoleResponse Response Object
type CreateDbRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDbRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbRoleResponse struct{}"
	}

	return strings.Join([]string{"CreateDbRoleResponse", string(data)}, " ")
}
