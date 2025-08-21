package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupMemberDetailDto struct {

	// **参数解释：** 唯一标识id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 资源id。
	SourceId *int32 `json:"source_id,omitempty"`

	// **参数解释：** 用户id。
	UserId *int32 `json:"user_id,omitempty"`

	// **参数解释：** 用户来源。 **取值范围：** 字符串长度不少于1，不超过1000。
	UserFrom *string `json:"user_from,omitempty"`

	// **参数解释：** 角色id。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleId *string `json:"role_id,omitempty"`

	// **参数解释：** 角色名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleName *string `json:"role_name,omitempty"`

	// **参数解释：** 角色中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	CnRoleName *string `json:"cn_role_name,omitempty"`

	// **参数解释：** 项目角色id。 **取值范围：** 字符串长度不少于1，不超过1000。
	ReqRoleId *string `json:"req_role_id,omitempty"`

	// **参数解释：** 项目角色名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	ReqRoleName *string `json:"req_role_name,omitempty"`

	// **参数解释：** 成员组id。 **取值范围：** 字符串长度不少于1，不超过1000。
	UserGroupId *string `json:"user_group_id,omitempty"`

	// **参数解释：** 代码组名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释：** 用户名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释：** 租户id。 **取值范围：** 字符串长度不少于1，不超过1000。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 租户名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释：** 用户昵称。 **取值范围：** 字符串长度不少于1，不超过1000。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 是否为代码组创建者。
	IsGroupCreator *bool `json:"is_group_creator,omitempty"`

	// **参数解释：** 是否为项目管理员。
	IsProjectAdmin *bool `json:"is_project_admin,omitempty"`

	// **参数解释：** 路径。 **取值范围：** 字符串长度不少于1，不超过1000。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 角色中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleChineseName *string `json:"role_chinese_name,omitempty"`

	// **参数解释：** 是否可移除。
	CanRemove *bool `json:"can_remove,omitempty"`

	// **参数解释：** 角色。
	AccessLevel *int32 `json:"access_level,omitempty"`

	// **参数解释：** 服务license状态。
	ServiceLicenseStatus *int32 `json:"service_license_status,omitempty"`

	// **参数解释：** 用户iam_id。 **取值范围：** 字符串长度不少于1，不超过1000。
	IamId *string `json:"iam_id,omitempty"`

	// **参数解释：** 是否为当前代码组成员。
	CurrentGroupMember *bool `json:"current_group_member,omitempty"`
}

func (o GroupMemberDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupMemberDetailDto struct{}"
	}

	return strings.Join([]string{"GroupMemberDetailDto", string(data)}, " ")
}
