package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RepositoryMemberDto 仓库成员详情
type RepositoryMemberDto struct {

	// 用户id
	UserId *int32 `json:"user_id,omitempty"`

	// 用户iamId
	UserIamId *string `json:"user_iam_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 用户昵称
	UserNickName *string `json:"user_nick_name,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`

	// 租户id
	TenantId *string `json:"tenant_id,omitempty"`

	// 是否为仓库所有者
	IsRepoCreator *int32 `json:"is_repo_creator,omitempty"`

	// 是否为父代码组所有者
	IsGroupCreator *int32 `json:"is_group_creator,omitempty"`

	// 是否为项目管理员
	IsProjectAdmin *int32 `json:"is_Project_admin,omitempty"`

	// 项目角色名称
	ProjectRoleName *string `json:"project_role_name,omitempty"`

	// 仓库角色名称
	RepositoryRoleName *string `json:"repository_role_name,omitempty"`

	// 仓库角色id
	RepositoryRoleId *string `json:"repository_role_Id,omitempty"`

	// 成员如果来自成员组，成员组名称
	MemberSource *string `json:"member_source,omitempty"`

	// 成员如果来自上层代码组，代码组名称
	MemberGroupSource *string `json:"member_group_source,omitempty"`

	// 成员来源id —— 代码组id或者成员组id
	MemberSourceId *string `json:"member_source_id,omitempty"`

	// 成员服务级权限状态： 1-使用中、0-已停用
	ServiceLicenseStatus *int32 `json:"service_license_status,omitempty"`

	// 是否有对应权限： true-有权限、false-无权限
	ActionEnabled *bool `json:"action_enabled,omitempty"`
}

func (o RepositoryMemberDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryMemberDto struct{}"
	}

	return strings.Join([]string{"RepositoryMemberDto", string(data)}, " ")
}
