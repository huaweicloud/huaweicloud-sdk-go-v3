package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersRequest Request Object
type ListUsersRequest struct {

	// 桌面用户名，长度范围为1-20，不能包含特殊字符，不能以数字开头。
	UserName *string `json:"user_name,omitempty"`

	// 用户邮箱。
	UserEmail *string `json:"user_email,omitempty"`

	// 用于分页查询，返回用户数量限制。如果不指定，则返回所有符合条件的用户。
	Limit *string `json:"limit,omitempty"`

	// 分页查询起始条数。
	Offset *string `json:"offset,omitempty"`

	// 用户描述查询，模糊匹配。
	Description *string `json:"description,omitempty"`

	// 激活类型，默认为用户激活。 * USER_ACTIVATE： 用户激活 * ADMIN_ACTIVATE： 管理员激活
	ActiveType *string `json:"active_type,omitempty"`

	// 桌面用户组名，精确匹配。
	GroupName *string `json:"group_name,omitempty"`

	// 用户是否已订阅协同，true/false
	ShareSpaceSubscription *bool `json:"share_space_subscription,omitempty"`

	// 用户是否已绑定协同桌面,true/false
	ShareSpaceDesktops *bool `json:"share_space_desktops,omitempty"`

	// 是否查询用户绑定的桌面数,true/false,默认true
	IsQueryTotalDesktops *bool `json:"is_query_total_desktops,omitempty"`
}

func (o ListUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersRequest struct{}"
	}

	return strings.Join([]string{"ListUsersRequest", string(data)}, " ")
}
