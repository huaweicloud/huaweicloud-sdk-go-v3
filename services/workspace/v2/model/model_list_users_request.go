package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListUsersRequest struct {

	// 桌面用户名，长度范围为1-20，不能包含特殊字符，不能以数字开头。
	UserName *string `json:"user_name,omitempty"`

	// 用于分页查询，返回用户数量限制。如果不指定，则返回所有符合条件的用户。
	Limit *string `json:"limit,omitempty"`

	// 分页查询起始条数。
	Offset *string `json:"offset,omitempty"`

	// 用户描述查询，模糊匹配。
	Description *string `json:"description,omitempty"`
}

func (o ListUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersRequest struct{}"
	}

	return strings.Join([]string{"ListUsersRequest", string(data)}, " ")
}
