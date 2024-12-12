package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersRequest Request Object
type ListUsersRequest struct {

	// 组织id，为空时查询所有用户。
	OrgId *string `json:"org_id,omitempty"`

	// 第几页。最小值：0
	Offset int32 `json:"offset"`

	// 每页多少条。最小值：10。最大值：100
	Limit int32 `json:"limit"`

	// 最长64位，用户名，支持模糊查询
	UserInfo *string `json:"user_info,omitempty"`
}

func (o ListUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersRequest struct{}"
	}

	return strings.Join([]string{"ListUsersRequest", string(data)}, " ")
}
