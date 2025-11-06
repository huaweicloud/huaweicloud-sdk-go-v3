package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManageableGroupsResponse Response Object
type ListManageableGroupsResponse struct {

	// 拥有管理权限的代码组列表
	Body           *[]ManageableGroupDto `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListManageableGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManageableGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListManageableGroupsResponse", string(data)}, " ")
}
