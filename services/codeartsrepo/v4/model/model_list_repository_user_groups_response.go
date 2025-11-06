package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryUserGroupsResponse Response Object
type ListRepositoryUserGroupsResponse struct {
	Body *[]RepositoryUserGroupDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryUserGroupsResponse", string(data)}, " ")
}
