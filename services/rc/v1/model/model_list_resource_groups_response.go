package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceGroupsResponse Response Object
type ListResourceGroupsResponse struct {
	ResourceGroups *[]CreateGroupsResponse `json:"resource_groups,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListResourceGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceGroupsResponse", string(data)}, " ")
}
