package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRoleResponse Response Object
type SetRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRoleResponse struct{}"
	}

	return strings.Join([]string{"SetRoleResponse", string(data)}, " ")
}
