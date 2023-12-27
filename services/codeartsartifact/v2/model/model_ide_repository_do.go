package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdeRepositoryDo struct {

	// 仓库名称
	RepositoryName *string `json:"repository_name,omitempty"`

	// 仓库类型
	Format *string `json:"format,omitempty"`

	// 仓库描述
	Description *string `json:"description,omitempty"`

	// release仓库名称
	Release *string `json:"release,omitempty"`

	// snapshot仓库名称
	Snapshot *string `json:"snapshot,omitempty"`

	// 路径
	IncludesPattern *string `json:"includes_pattern,omitempty"`

	// 共享权限级别
	ShareRight *string `json:"share_right,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 仓库类别，本地仓或聚合仓
	Type *string `json:"type,omitempty"`
}

func (o IdeRepositoryDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdeRepositoryDo struct{}"
	}

	return strings.Join([]string{"IdeRepositoryDo", string(data)}, " ")
}
