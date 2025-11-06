package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryDirDto struct {

	// **参数解释：** 目录路径。 **约束限制：** 目录路径层级最大不能超过29。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	FilePath string `json:"file_path"`

	// **参数解释：** 分支名。 **取值范围：** 最小1个字节，最大200字节 **约束限制：** 该仓库下的已有分支。
	BranchName string `json:"branch_name"`

	// **参数解释：** 提交信息。 **取值范围：** 不涉及。
	CommitMessage string `json:"commit_message"`
}

func (o RepositoryDirDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryDirDto struct{}"
	}

	return strings.Join([]string{"RepositoryDirDto", string(data)}, " ")
}
