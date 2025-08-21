package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubgroupAndProjectBaseDto A empty base object for subgroup and project.
type SubgroupAndProjectBaseDto struct {

	// **参数解释：** 项目id。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 项目名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 角色中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleNamecn *string `json:"role_namecn,omitempty"`

	// **参数解释：** 角色英文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleNameen *string `json:"role_nameen,omitempty"`

	// **参数解释：** 全名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	FullName *string `json:"full_name,omitempty"`

	// **参数解释：** 全路径。 **取值范围：** 字符串长度不少于1，不超过1000。
	FullPath *string `json:"full_path,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间戳。 **取值范围：** 字符串长度不少于1，不超过1000。
	UpdatedAtTimestamp *string `json:"updated_at_timestamp,omitempty"`

	// **参数解释：** 开始时间戳。 **取值范围：** 字符串长度不少于1，不超过1000。
	StarTime *string `json:"star_time,omitempty"`

	// **参数解释：** 是否收藏。
	Starred *bool `json:"starred,omitempty"`

	// **参数解释：** 开发模式，cr,\"normal\"。 **取值范围：** 字符串长度不少于1，不超过1000。
	DevelopMode *string `json:"develop_mode,omitempty"`

	// **参数解释：** 仓库或者代码组id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 代码组或仓库名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 路径。 **取值范围：** 字符串长度不少于1，不超过1000。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 代码组层级。
	GroupLevel *int32 `json:"group_level,omitempty"`

	// **参数解释：** 代码组或仓库描述。 **取值范围：** 字符串长度不少于1，不超过1000。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 子代码组数量。
	SubgroupCount *int32 `json:"subgroup_count,omitempty"`

	// **参数解释：** 仓库数量。
	ProjectCount *int32 `json:"project_count,omitempty"`

	// **参数解释：** 代码组角色。
	GroupRole *int32 `json:"group_role,omitempty"`

	// **参数解释：** 代码组成员数量。
	GroupMembersCount *int32 `json:"group_members_count,omitempty"`

	// **参数解释：** 资源类型 group,project。 **取值范围：** 字符串长度不少于1，不超过1000。
	DescendantType *string `json:"descendant_type,omitempty"`

	// **参数解释：** 可见性level 0(私有),20(公开)
	VisibilityLevel *int32 `json:"visibility_level,omitempty"`

	// **参数解释：** 可见性 private public。 **取值范围：** 字符串长度不少于1，不超过1000。
	Visibility *string `json:"visibility,omitempty"`

	// **参数解释：** 当前用户是否为项目创建者。
	IsProjectAdmin *int32 `json:"is_project_admin,omitempty"`

	// **参数解释：** 当前用户是否为代码组创建者。
	IsGroupCreator *int32 `json:"is_group_creator,omitempty"`

	// **参数解释：** 当前用户是否为仓库创建者。
	IsRepoCreator *int32 `json:"is_repo_creator,omitempty"`

	// **参数解释：** 角色展示标记。
	RoleShowFlag *int32 `json:"role_show_flag,omitempty"`

	// **参数解释：** 仓库的uuid。 **取值范围：** 字符串长度不少于1，不超过1000。
	Uuid *string `json:"uuid,omitempty"`

	// **参数解释：** fork数量。
	ForksCount *int32 `json:"forks_count,omitempty"`

	// **参数解释：** 是否为kia。
	IsKia *bool `json:"is_kia,omitempty"`

	// **参数解释：** 是否为所有者。
	IsOwner *bool `json:"is_owner,omitempty"`

	// **参数解释：** 是否为存档。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 仓库的最后更新时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	LastRepositoryUpdatedAt *string `json:"last_repository_updated_at,omitempty"`

	// **参数解释：** 开启的mr数量。
	OpenMergeRequestsCount *int32 `json:"open_merge_requests_count,omitempty"`

	// **参数解释：** 所有的mr数量。
	AllMergeRequestsCount *int32 `json:"all_merge_requests_count,omitempty"`

	// **参数解释：** 仓库角色。
	ProjectRole *int32 `json:"project_role,omitempty"`

	// **参数解释：** fork数量。
	ProjectMembersCount *int32 `json:"project_members_count,omitempty"`

	ProjectCreator *ProjectCreatorDto `json:"project_creator,omitempty"`

	// **参数解释：** fork数量。
	StarCount *int32 `json:"star_count,omitempty"`

	// **参数解释：** tag列表。
	TagList *[]string `json:"tag_list,omitempty"`

	// **参数解释：** 仓库的http url。 **取值范围：** 字符串长度不少于1，不超过1000。
	HttpUrlToRepo *string `json:"http_url_to_repo,omitempty"`

	// **参数解释：** 仓库的ssh url。 **取值范围：** 字符串长度不少于1，不超过1000。
	SshUrlToRepo *string `json:"ssh_url_to_repo,omitempty"`

	// **参数解释：** 状态。
	Status *int32 `json:"status,omitempty"`

	// **参数解释：** 活跃统计。
	ActiveStatistic *[]int32 `json:"active_statistic,omitempty"`

	// **参数解释：** 安全标签。 **取值范围：** 字符串长度不少于1，不超过1000。
	SecurityTag *string `json:"security_tag,omitempty"`
}

func (o SubgroupAndProjectBaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubgroupAndProjectBaseDto struct{}"
	}

	return strings.Join([]string{"SubgroupAndProjectBaseDto", string(data)}, " ")
}
