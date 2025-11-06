package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubmoduleReqDto struct {

	// **参数解释：** 分支名。 **取值范围：** 最小1个字节，最大200字节 **约束限制：** 该仓库下的已有分支。
	BranchName string `json:"branch_name"`

	// 子模块在该仓库下的文件路径
	FilePath string `json:"file_path"`

	// 子模块所在仓库的仓库ID
	SubrepoId string `json:"subrepo_id"`

	// 提交信息
	CommitMessage string `json:"commit_message"`

	// **参数解释：** 子模块分支名。 **取值范围：** 最小1个字节，最大200字节 **约束限制：** 目标仓库下的已有分支。
	SubrepoBranch string `json:"subrepo_branch"`
}

func (o SubmoduleReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubmoduleReqDto struct{}"
	}

	return strings.Join([]string{"SubmoduleReqDto", string(data)}, " ")
}
