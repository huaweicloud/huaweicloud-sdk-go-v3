package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdeRepositoryPair struct {

	// 仓库名称
	RepoName *string `json:"repo_name,omitempty"`

	// 路径
	IncludesPattern *string `json:"includes_pattern,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// snapshot仓库名称
	Snapshot *string `json:"snapshot,omitempty"`

	// release仓库名称
	Release *string `json:"release,omitempty"`
}

func (o IdeRepositoryPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdeRepositoryPair struct{}"
	}

	return strings.Join([]string{"IdeRepositoryPair", string(data)}, " ")
}
