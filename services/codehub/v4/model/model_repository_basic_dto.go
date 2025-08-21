package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepositoryBasicDto struct {

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

	// **参数解释：** 仓库页面链接。
	WebUrl *string `json:"web_url,omitempty"`

	// **参数解释：** 仓库readme文件链接。
	ReadmeUrl *string `json:"readme_url,omitempty"`

	// **参数解释：** 仓库所属项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 仓库所属项目名称。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 仓库开发模式。 **取值范围：** - normal - CR
	DevelopMode *RepositoryBasicDtoDevelopMode `json:"develop_mode,omitempty"`

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
}

func (o RepositoryBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryBasicDto struct{}"
	}

	return strings.Join([]string{"RepositoryBasicDto", string(data)}, " ")
}

type RepositoryBasicDtoDevelopMode struct {
	value string
}

type RepositoryBasicDtoDevelopModeEnum struct {
	NORMAL RepositoryBasicDtoDevelopMode
	CR     RepositoryBasicDtoDevelopMode
}

func GetRepositoryBasicDtoDevelopModeEnum() RepositoryBasicDtoDevelopModeEnum {
	return RepositoryBasicDtoDevelopModeEnum{
		NORMAL: RepositoryBasicDtoDevelopMode{
			value: "normal",
		},
		CR: RepositoryBasicDtoDevelopMode{
			value: "CR",
		},
	}
}

func (c RepositoryBasicDtoDevelopMode) Value() string {
	return c.value
}

func (c RepositoryBasicDtoDevelopMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryBasicDtoDevelopMode) UnmarshalJSON(b []byte) error {
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
