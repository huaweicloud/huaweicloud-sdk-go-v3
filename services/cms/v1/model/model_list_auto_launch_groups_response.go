package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoLaunchGroupsResponse Response Object
type ListAutoLaunchGroupsResponse struct {

	// 智能购买组列表
	AutoLaunchGroups *[]AutoLaunchGroupInfo `json:"auto_launch_groups,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAutoLaunchGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchGroupsResponse", string(data)}, " ")
}
