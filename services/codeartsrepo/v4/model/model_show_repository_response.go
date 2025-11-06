package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepositoryResponse Response Object
type ShowRepositoryResponse struct {

	// **参数解释：** 仓库ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库描述信息。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 仓库名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 仓库完整名称。
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`

	// **参数解释：** 仓库路径。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 仓库完整路径。
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 是否归档。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 仓库ssh地址。
	SshUrlToRepo *string `json:"ssh_url_to_repo,omitempty"`

	// **参数解释：** 仓库http地址。
	HttpUrlToRepo *string `json:"http_url_to_repo,omitempty"`

	// **参数解释：** 仓库所属项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 仓库所属项目名称。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 仓库开发模式。 **取值范围：** - normal - CR
	DevelopMode *ShowRepositoryResponseDevelopMode `json:"develop_mode,omitempty"`

	// **参数解释：** 审核状态。
	ModerationResult *bool `json:"moderation_result,omitempty"`

	// **参数解释：** 仓库默认分支 **约束限制：** 不涉及。
	DefaultBranch *string `json:"default_branch,omitempty"`

	// **参数解释：** 仓库图标url **约束限制：** 不涉及。
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// **参数解释：** 关注数 **约束限制：** 不涉及。
	StarCount *int32 `json:"star_count,omitempty"`

	// **参数解释：** fork数 **约束限制：** 不涉及。
	ForksCount *int32 `json:"forks_count,omitempty"`

	// **参数解释：** 开启issue数 **约束限制：** 不涉及。
	OpenIssuesCount *int32 `json:"open_issues_count,omitempty"`

	// **参数解释：** 开启中的CR、MR数量 **约束限制：** 不涉及。
	OpenMergeRequestsCount *int32 `json:"open_merge_requests_count,omitempty"`

	// **参数解释：** 最后活跃时间 **约束限制：** 不涉及。
	LastActivityAt *string `json:"last_activity_at,omitempty"`

	Namespace *NamespaceBasicDto `json:"namespace,omitempty"`

	// **参数解释：** 空仓库 **约束限制：** 不涉及。
	EmptyRepo *bool `json:"empty_repo,omitempty"`

	// **参数解释：** 是否已关注 **约束限制：** 不涉及。
	Starred *bool `json:"starred,omitempty"`

	// **参数解释：** 仓库可见等级 **约束限制：** 不涉及。
	Visibility *string `json:"visibility,omitempty"`

	// **参数解释：** 仓库保密等级 **约束限制：** 不涉及。
	SecurityTag *string `json:"security_tag,omitempty"`

	// **参数解释：** 仓库保密等级 **约束限制：** 不涉及。
	Security *string `json:"security,omitempty"`

	// **参数解释：** 网络类型 **约束限制：** 不涉及。
	NetworkType *string `json:"network_type,omitempty"`

	Owner *RepositoryUserBasicDto `json:"owner,omitempty"`

	Creator *RepositoryUserBasicDto `json:"creator,omitempty"`

	// **参数解释：** 创建者ID **约束限制：** 不涉及。
	CreatorId *int32 `json:"creator_id,omitempty"`

	ForkedFromRepository *RepositorySimpleDto `json:"forked_from_repository,omitempty"`

	// **参数解释：** 仓库唯一标识符。 **约束限制：** 不涉及。
	Uuid *string `json:"uuid,omitempty"`

	// **参数解释：** 祖先仓库ID列表。 **约束限制：** 不涉及。
	AncestorIds *[]int32 `json:"ancestor_ids,omitempty"`

	// **参数解释：** 祖先仓库名称列表。 **约束限制：** 不涉及。
	AncestorNames *[]string `json:"ancestor_names,omitempty"`

	// **参数解释：** 导入状态。 **约束限制：** 不涉及。
	ImportStatus *string `json:"import_status,omitempty"`

	// **参数解释：** 导入源地址。 **约束限制：** 不涉及。
	ImportUrl *string `json:"import_url,omitempty"`

	// **参数解释：** 导入错误信息。 **约束限制：** 不涉及。
	ImportError *string `json:"import_error,omitempty"`

	// **参数解释：** 仓库类型。 **约束限制：** 不涉及。
	RepoType *string `json:"repo_type,omitempty"`

	// **参数解释：** 是否仅在流水线成功时允许合并。 **约束限制：** 不涉及。
	OnlyAllowMergeIfPipelineSucceeds *bool `json:"only_allow_merge_if_pipeline_succeeds,omitempty"`

	// **参数解释：** 是否启用访问请求。 **约束限制：** 不涉及。
	RequestAccessEnabled *bool `json:"request_access_enabled,omitempty"`

	// **参数解释：** 是否仅在所有检视意见解决时允许合并。 **约束限制：** 不涉及。
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`

	// **参数解释：** 合并方法。 **约束限制：** 不涉及。
	MergeMethod *string `json:"merge_method,omitempty"`

	// **参数解释：** fork关联仓库列表。 **约束限制：** 不涉及。
	ForkNetworkRepositories *[]RepositoryIdentityDto `json:"fork_network_repositories,omitempty"`

	Permissions *PermissionsDto `json:"permissions,omitempty"`

	// **参数解释：** 仓库类型。 **约束限制：** 不涉及。
	RepositoryType *string `json:"repository_type,omitempty"`

	Statistics *RepositoryStatisticsDto `json:"statistics,omitempty"`

	// **参数解释：** 分支数量。 **约束限制：** 不涉及。
	BranchCount *int32 `json:"branch_count,omitempty"`

	// **参数解释：** Tag数量。 **约束限制：** 不涉及。
	TagCount *int32 `json:"tag_count,omitempty"`

	// **参数解释：** 标签数量。 **约束限制：** 不涉及。
	LabelCount *int32 `json:"label_count,omitempty"`

	// **参数解释：** 成员数量。 **约束限制：** 不涉及。
	MemberCount    *int32 `json:"member_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryResponse", string(data)}, " ")
}

type ShowRepositoryResponseDevelopMode struct {
	value string
}

type ShowRepositoryResponseDevelopModeEnum struct {
	NORMAL ShowRepositoryResponseDevelopMode
	CR     ShowRepositoryResponseDevelopMode
}

func GetShowRepositoryResponseDevelopModeEnum() ShowRepositoryResponseDevelopModeEnum {
	return ShowRepositoryResponseDevelopModeEnum{
		NORMAL: ShowRepositoryResponseDevelopMode{
			value: "normal",
		},
		CR: ShowRepositoryResponseDevelopMode{
			value: "CR",
		},
	}
}

func (c ShowRepositoryResponseDevelopMode) Value() string {
	return c.value
}

func (c ShowRepositoryResponseDevelopMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepositoryResponseDevelopMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
