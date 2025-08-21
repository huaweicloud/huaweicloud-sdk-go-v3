package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositorySimplestDto struct {

	// **参数解释：** 仓库ID **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库名称。 **约束限制：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 命名空间。 **约束限制：** 不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释：** 仓库路径。 **约束限制：** 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 开发模式。 **约束限制：** 不涉及。
	DevelopMode *string `json:"develop_mode,omitempty"`

	// **参数解释：** 可见性。 **约束限制：** 不涉及。
	Visibility *string `json:"visibility,omitempty"`

	// **参数解释：** 安全级别。 **约束限制：** 不涉及。
	Security *string `json:"security,omitempty"`

	// **参数解释：** 关注数。 **约束限制：** 不涉及。
	StarCount *int32 `json:"star_count,omitempty"`

	// **参数解释：** Fork数。 **约束限制：** 不涉及。
	ForksCount *int32 `json:"forks_count,omitempty"`

	// **参数解释：** 开放的问题数。 **约束限制：** 不涉及。
	OpenIssuesCount *int32 `json:"open_issues_count,omitempty"`

	// **参数解释：** 开放的合并请求数。 **约束限制：** 不涉及。
	OpenMergeRequestsCount *int32 `json:"open_merge_requests_count,omitempty"`

	// **参数解释：** 是否已关注。 **约束限制：** 不涉及。
	Starred *bool `json:"starred,omitempty"`

	// **参数解释：** 包含命名空间的仓库名称。 **约束限制：** 不涉及。
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`

	// **参数解释：** 最后活跃时间。 **约束限制：** 不涉及。
	LastActivityAt *string `json:"last_activity_at,omitempty"`

	ForkedFromRepository *ForkedFromRepositorySimplestDto `json:"forked_from_repository,omitempty"`

	// **参数解释：** 权限。 **约束限制：** 不涉及。
	Permissions *int32 `json:"permissions,omitempty"`

	// **参数解释：** 是否归档。 **约束限制：** 不涉及。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 成员数。 **约束限制：** 不涉及。
	MemberCount *int32 `json:"member_count,omitempty"`

	// **参数解释：** 仓库唯一标识符。 **约束限制：** 不涉及。
	Uuid *string `json:"uuid,omitempty"`

	// **参数解释：** 仓库描述。 **约束限制：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 最后更新时间。 **约束限制：** 不涉及。
	LastRepositoryUpdatedAt *string `json:"last_repository_updated_at,omitempty"`

	// **参数解释：** SSH仓库URL。 **约束限制：** 不涉及。
	SshUrlToRepo *string `json:"ssh_url_to_repo,omitempty"`

	// **参数解释：** HTTP仓库URL。 **约束限制：** 不涉及。
	HttpUrlToRepo *string `json:"http_url_to_repo,omitempty"`

	// **参数解释：** 仓库状态。 **约束限制：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 活跃统计。 **约束限制：** 不涉及。
	ActiveStatistic *[]int32 `json:"active_statistic,omitempty"`

	// **参数解释：** 项目名称。 **约束限制：** 不涉及。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 项目ID。 **约束限制：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 项目是否删除。 **约束限制：** 不涉及。
	ProjectIsDelete *bool `json:"project_is_delete,omitempty"`

	// **参数解释：** 创建者ID **约束限制：** 不涉及。
	CreatorId *int32 `json:"creator_id,omitempty"`
}

func (o RepositorySimplestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositorySimplestDto struct{}"
	}

	return strings.Join([]string{"RepositorySimplestDto", string(data)}, " ")
}
