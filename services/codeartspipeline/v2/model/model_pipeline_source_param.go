package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineSourceParam **参数解释**： 流水线源参数。 **取值范围**： 不涉及。
type PipelineSourceParam struct {

	// **参数解释**： 代码/制品源参数 - 代码仓/制品源别名。 **取值范围**： 别名仅支持输入大小写英文字母、数字、“_”，至多128个字符。
	Alias *string `json:"alias,omitempty"`

	// **参数解释**： git代码仓类型。 **取值范围**： - codehub。 - gitee。 - github。 - gitcode。 - gitlab。
	GitType *string `json:"git_type,omitempty"`

	// **参数解释**： Repo代码仓ID。 **取值范围**： 不涉及。
	CodehubId *string `json:"codehub_id,omitempty"`

	// **参数解释**： 扩展点id。 **取值范围**： 不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**： 默认分支。 **取值范围**： 不涉及。
	DefaultBranch *string `json:"default_branch,omitempty"`

	// **参数解释**： git链接。 **取值范围**： 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释**： ssh_git链接。 **取值范围**： 不涉及。
	SshGitUrl *string `json:"ssh_git_url,omitempty"`

	// **参数解释**： 网页url。 **取值范围**： 不涉及。
	WebUrl *string `json:"web_url,omitempty"`

	// **参数解释**： 流水线源名称。 **取值范围**： 不涉及。
	RepoName *string `json:"repo_name,omitempty"`

	// **参数解释**： 制品源类型。 **取值范围**： 仅包含[generic，docker]。
	ArtifactType *string `json:"artifact_type,omitempty"`

	// **参数解释**： 制品源类型名。 **取值范围**： 不涉及。
	ArtifactTypeName *string `json:"artifact_type_name,omitempty"`

	// **参数解释**： 过滤分支。 **取值范围**： 不涉及。
	BranchFilter *string `json:"branch_filter,omitempty"`

	// **参数解释**： 制品源所在目录。 **取值范围**： 不涉及。
	Directory *string `json:"directory,omitempty"`

	// **参数解释**： 目录ID。 **取值范围**： 不涉及。
	DirectoryId *string `json:"directory_id,omitempty"`

	// **参数解释**： Docker组织。 **取值范围**： 不涉及。
	Organization *string `json:"organization,omitempty"`

	// **参数解释**： 软件包名称。 **取值范围**： 不涉及。
	PackageName *string `json:"package_name,omitempty"`

	// **参数解释**： 制品源版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 获取制品源版本的策略。 **取值范围**： 仅包含[latest，specificVersion]。
	VersionStrategy *string `json:"version_strategy,omitempty"`

	// **参数解释**： 制品源名称。 **取值范围**： 不涉及。
	SourceSystem *string `json:"source_system,omitempty"`
}

func (o PipelineSourceParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineSourceParam struct{}"
	}

	return strings.Join([]string{"PipelineSourceParam", string(data)}, " ")
}
