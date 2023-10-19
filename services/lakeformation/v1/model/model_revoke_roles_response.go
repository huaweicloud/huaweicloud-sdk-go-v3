package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokeRolesResponse Response Object
type RevokeRolesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RevokeRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeRolesResponse struct{}"
	}

	return strings.Join([]string{"RevokeRolesResponse", string(data)}, " ")
}
