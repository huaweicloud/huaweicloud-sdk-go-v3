package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigWorkspaceUserbody 工作空间用户列表
type ApigWorkspaceUserbody struct {

	// 记录id
	Id *string `json:"id,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 租户名
	DomainName *string `json:"domain_name,omitempty"`

	// 租户名
	DisplayUserName *string `json:"display_user_name,omitempty"`

	// 是否是空间所有者
	DomainOwner *bool `json:"domain_owner,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 角色列表
	Roles *[]ApigRoleVo `json:"roles,omitempty"`

	// 创建时间
	CreateTime float32 `json:"create_time,omitempty"`

	// 创建人员
	CreateUser *string `json:"create_user,omitempty"`

	// 更新时间
	UpdateTime float32 `json:"update_time,omitempty"`

	// 更新人员
	UpdateUser *string `json:"update_user,omitempty"`

	// 用户类型，0用户，1用户组
	Type *int32 `json:"type,omitempty"`
}

func (o ApigWorkspaceUserbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigWorkspaceUserbody struct{}"
	}

	return strings.Join([]string{"ApigWorkspaceUserbody", string(data)}, " ")
}
