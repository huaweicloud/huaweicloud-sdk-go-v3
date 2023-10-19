package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateRolesResponse Response Object
type AssociateRolesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRolesResponse struct{}"
	}

	return strings.Join([]string{"AssociateRolesResponse", string(data)}, " ")
}
