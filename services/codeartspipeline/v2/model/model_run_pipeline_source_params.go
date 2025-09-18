package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineSourceParams **参数解释**： 代码源相关参数。 **取值范围**： 不涉及。
type RunPipelineSourceParams struct {

	// **参数解释**： 代码仓类型。 **取值范围**： - codehub。 - gitee。 - github。 - gitcode。 - gitlab。
	GitType *string `json:"git_type,omitempty"`

	// **参数解释**： 代码仓https地址。 **取值范围**： 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释**： 代码仓ssh地址。 **取值范围**： 不涉及。
	SshGitUrl *string `json:"ssh_git_url,omitempty"`

	// **参数解释**： 代码仓页面地址。 **取值范围**： 不涉及。
	WebUrl *string `json:"web_url,omitempty"`

	// **参数解释**： 代码仓名。 **取值范围**： 不涉及。
	RepoName *string `json:"repo_name,omitempty"`

	// **参数解释**： 默认分支。 **取值范围**： 不涉及。
	DefaultBranch *string `json:"default_branch,omitempty"`

	// **参数解释**： 扩展点ID。 **取值范围**： 不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**： Repo代码仓ID。 **取值范围**： 不涉及。
	CodehubId *string `json:"codehub_id,omitempty"`

	// **参数解释**： 代码仓别名。 **取值范围**： 不涉及。
	Alias *string `json:"alias,omitempty"`

	BuildParams *RunPipelineSourceParamsBuildParams `json:"build_params,omitempty"`
}

func (o RunPipelineSourceParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineSourceParams struct{}"
	}

	return strings.Join([]string{"RunPipelineSourceParams", string(data)}, " ")
}
