package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRoleResponse Response Object
type DeleteRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRoleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRoleResponse", string(data)}, " ")
}
