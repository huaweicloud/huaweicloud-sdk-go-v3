package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupsResponse Response Object
type ListGroupsResponse struct {

	// 满足查询条件的用户组对象列表
	Groups *[]GroupDto `json:"groups,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupsResponse", string(data)}, " ")
}
