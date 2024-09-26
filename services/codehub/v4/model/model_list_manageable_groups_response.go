package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManageableGroupsResponse Response Object
type ListManageableGroupsResponse struct {

	// 仓库ip白名单列表
	Body *[]GroupsManageableDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListManageableGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManageableGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListManageableGroupsResponse", string(data)}, " ")
}
