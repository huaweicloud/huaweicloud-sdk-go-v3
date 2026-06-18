package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepoStatisticsSummaryResponse Response Object
type ShowRepoStatisticsSummaryResponse struct {

	// **参数解释：** 仓库名称。 **取值范围：** 最小1个字节，最大200字节 **默认取值：** 不涉及。
	RepoName *string `json:"repo_name,omitempty"`

	// **参数解释：** 默认分支的提交数量。 **取值范围：** 最小0 **默认取值：** 0
	CommitCount *int32 `json:"commit_count,omitempty"`

	// **参数解释：** 仓库占用磁盘空间大小。 **取值范围：** 最小0 **默认取值：** 0
	RepoSize *string `json:"repo_size,omitempty"`

	// **参数解释：** 仓库最新的提交日期，格式yyyy-MM-dd'T'HH:mm:ssXXX,例：2025-10-30T08:27:43.000Z **取值范围：** 不涉及。 **默认取值：** 不涉及。
	LastCommitTime *string `json:"last_commit_time,omitempty"`

	// **参数解释：** 默认分支的代码行数。 **取值范围：** 最小0 **默认取值：** 0。
	CodeLines *int32 `json:"code_lines,omitempty"`

	// **参数解释：** 仓库分支数量 **取值范围：** 最小0 **默认取值：** 0
	BranchCount    *int32 `json:"branch_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRepoStatisticsSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoStatisticsSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowRepoStatisticsSummaryResponse", string(data)}, " ")
}
