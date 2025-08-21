package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupBaseDto A empty base object for group.
type GroupBaseDto struct {

	// **参数解释：** 项目id。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 项目名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 代码组id。
	AncestorIds *[]int32 `json:"ancestor_ids,omitempty"`

	// **参数解释：** 代码组名称。
	AncestorNames *[]string `json:"ancestor_names,omitempty"`

	// **参数解释：** 开发模式，normal，cr。 **取值范围：** 字符串长度不少于1，不超过1000。
	DevelopMode *string `json:"develop_mode,omitempty"`

	// **参数解释：** 记录id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 路径。 **取值范围：** 字符串长度不少于1，不超过1000。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 代码组层级。
	GroupLevel *int32 `json:"group_level,omitempty"`

	// **参数解释：** 描述。 **取值范围：** 字符串长度不少于1，不超过1000。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 子代码组数量。
	SubgroupCount *int32 `json:"subgroup_count,omitempty"`

	// **参数解释：** 仓库数量。
	ProjectCount *int32 `json:"project_count,omitempty"`

	// **参数解释：** 代码组角色。
	GroupRole *int32 `json:"group_role,omitempty"`

	// **参数解释：** 代码组成员数量。
	GroupMembersCount *int32 `json:"group_members_count,omitempty"`

	// **参数解释：** 类型。 **取值范围：** 字符串长度不少于1，不超过1000。
	DescendantType *string `json:"descendant_type,omitempty"`

	// **参数解释：** 可见性 0 20。
	VisibilityLevel *int32 `json:"visibility_level,omitempty"`

	// **参数解释：** 可见性 private public。 **取值范围：** 字符串长度不少于1，不超过1000。
	Visibility *string `json:"visibility,omitempty"`

	// **参数解释：** 是否为项目创建者。
	IsProjectAdmin *int32 `json:"is_project_admin,omitempty"`

	// **参数解释：** 是否为代码组创建者。
	IsGroupCreator *int32 `json:"is_group_creator,omitempty"`

	// **参数解释：** 是否为仓库创建者。
	IsRepoCreator *int32 `json:"is_repo_creator,omitempty"`

	// **参数解释：** lfs是否开启。
	LfsEnabled *bool `json:"lfs_enabled,omitempty"`

	// **参数解释：** 全名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	FullName *string `json:"full_name,omitempty"`

	// **参数解释：** 全路径。 **取值范围：** 字符串长度不少于1，不超过1000。
	FullPath *string `json:"full_path,omitempty"`

	// **参数解释：** item类型。 **取值范围：** 字符串长度不少于1，不超过1000。
	ItemType *string `json:"item_type,omitempty"`

	// **参数解释：** 父代码组id。
	ParentId *int32 `json:"parent_id,omitempty"`

	MyRole *GroupMyRoleDtoV4 `json:"my_role,omitempty"`

	// **参数解释：** 成员。
	Members *int32 `json:"members,omitempty"`

	// **参数解释：** url地址。 **取值范围：** 字符串长度不少于1，不超过1000。
	WebUrl *string `json:"web_url,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 子代码组数量。
	SubGroupCount *int32 `json:"sub_group_count,omitempty"`

	// **参数解释：** 是否为最后所有者。
	LastOwner *bool `json:"last_owner,omitempty"`

	// **参数解释：** 是否关注。
	Starred *bool `json:"starred,omitempty"`
}

func (o GroupBaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupBaseDto struct{}"
	}

	return strings.Join([]string{"GroupBaseDto", string(data)}, " ")
}
