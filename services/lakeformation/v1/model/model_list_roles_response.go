package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRolesResponse Response Object
type ListRolesResponse struct {
	Roles *[]Role `json:"roles,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRolesResponse struct{}"
	}

	return strings.Join([]string{"ListRolesResponse", string(data)}, " ")
}
