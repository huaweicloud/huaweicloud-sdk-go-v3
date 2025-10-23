package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupProtectedRefsUserGroupsRequest Request Object
type ListGroupProtectedRefsUserGroupsRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 查询关键字，可模糊匹配用户名称、用户昵称、租户名称。
	Search *string `json:"search,omitempty"`
}

func (o ListGroupProtectedRefsUserGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupProtectedRefsUserGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListGroupProtectedRefsUserGroupsRequest", string(data)}, " ")
}
