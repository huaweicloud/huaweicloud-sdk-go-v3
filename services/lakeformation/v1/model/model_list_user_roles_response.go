package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserRolesResponse Response Object
type ListUserRolesResponse struct {
	Roles *[]Role `json:"roles,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListUserRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserRolesResponse struct{}"
	}

	return strings.Join([]string{"ListUserRolesResponse", string(data)}, " ")
}
