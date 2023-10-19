package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRolesResponse Response Object
type UpdateRolesResponse struct {
	Body           *[]Role `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRolesResponse struct{}"
	}

	return strings.Join([]string{"UpdateRolesResponse", string(data)}, " ")
}
