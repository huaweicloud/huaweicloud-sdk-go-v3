package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubmoduleDto struct {

	// **参数解释：** 仓库ID。
	RepoId *int32 `json:"repo_id,omitempty"`

	// **参数解释：** 分支名。 **取值范围：** 最小1个字节，最大200字节。
	Branch *string `json:"branch,omitempty"`

	// **参数解释：** 分支名。 **取值范围：** 最小1个字节，最大200字节
	Path *string `json:"path,omitempty"`

	// **参数解释：** 子模块Git地址。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释：** 子模块分支名。 **取值范围：** 最小1个字节，最大200字节。
	SubmoduleBranch *string `json:"submodule_branch,omitempty"`

	// 组织名/组织名.../仓库名
	NamespaceUuid *string `json:"namespace_uuid,omitempty"`

	// **参数解释：** 子模块仓库ID。
	SubmoduleRepoId *int32 `json:"submodule_repo_id,omitempty"`

	// **参数解释：** 子模块仓库名称。
	RepoName *string `json:"repo_name,omitempty"`

	// **参数解释：** 子模块仓库提交。
	SubCommitId *string `json:"sub_commitId,omitempty"`

	// **参数解释：** 部署秘钥同步状态。 **取值范围：** - 0，不同步。 - 1，同步。
	DeployKeyStatus *int32 `json:"deployKey_status,omitempty"`

	// **参数解释：** 子模块状态。 **取值范围：** - 0，异常。 - 1，正常。
	Status *int32 `json:"status,omitempty"`
}

func (o SubmoduleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubmoduleDto struct{}"
	}

	return strings.Join([]string{"SubmoduleDto", string(data)}, " ")
}
