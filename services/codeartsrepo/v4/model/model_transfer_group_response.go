package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferGroupResponse Response Object
type TransferGroupResponse struct {

	// 代码组id
	Id *int32 `json:"id,omitempty"`

	// 代码组全名称
	FullName *string `json:"full_name,omitempty"`

	// 代码组全路径
	FullPath *string `json:"full_path,omitempty"`

	MyRole *GroupMyRoleDto `json:"my_role,omitempty"`

	// 代码组名称
	Name *string `json:"name,omitempty"`

	// 代码组父级id
	ParentId *int32 `json:"parent_id,omitempty"`

	// 代码组所有者id
	CreatorId      *int32 `json:"creator_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o TransferGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferGroupResponse struct{}"
	}

	return strings.Join([]string{"TransferGroupResponse", string(data)}, " ")
}
