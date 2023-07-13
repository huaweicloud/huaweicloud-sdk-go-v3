package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupsForDomainResponse Response Object
type ListGroupsForDomainResponse struct {
	UserGroup *[]UserGroup `json:"user_group,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListGroupsForDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsForDomainResponse struct{}"
	}

	return strings.Join([]string{"ListGroupsForDomainResponse", string(data)}, " ")
}
