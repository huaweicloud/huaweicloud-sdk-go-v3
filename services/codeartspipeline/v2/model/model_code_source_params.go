package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeSourceParams **参数解释**： 流水线源参数，包含流水线源的详细信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type CodeSourceParams struct {

	// **参数解释**： 代码仓类型。 **约束限制**： 不涉及。 **取值范围**： - codehub。 - gitee。 - github。 - gitcode。 - gitlab。 **默认取值**： 不涉及。
	GitType *string `json:"git_type,omitempty"`

	// **参数解释**： CodeArts Repo代码仓ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CodehubId *string `json:"codehub_id,omitempty"`

	// **参数解释**： 代码源扩展点ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**： 默认分支。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DefaultBranch *string `json:"default_branch,omitempty"`

	// **参数解释**： Git仓库https地址，例如https://example.com/CloudPipelinezycs00001/2000.git。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释**： ssh_git链接地址，例如https://example.com/CloudPipelinezycs00001/2000.git。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SshGitUrl *string `json:"ssh_git_url,omitempty"`

	// **参数解释**： 网页url。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WebUrl *string `json:"web_url,omitempty"`

	// **参数解释**： 流水线源名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RepoName *string `json:"repo_name,omitempty"`

	// **参数解释**： 代码仓别名。 **约束限制**： 不涉及。 **取值范围**： 仅支持输入大小写英文字母、数字、“_”，至多128个字符。 **默认取值**： 不涉及。
	Alias *string `json:"alias,omitempty"`
}

func (o CodeSourceParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeSourceParams struct{}"
	}

	return strings.Join([]string{"CodeSourceParams", string(data)}, " ")
}
