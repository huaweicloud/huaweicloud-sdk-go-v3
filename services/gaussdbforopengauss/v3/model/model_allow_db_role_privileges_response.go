package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowDbRolePrivilegesResponse Response Object
type AllowDbRolePrivilegesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AllowDbRolePrivilegesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowDbRolePrivilegesResponse struct{}"
	}

	return strings.Join([]string{"AllowDbRolePrivilegesResponse", string(data)}, " ")
}
