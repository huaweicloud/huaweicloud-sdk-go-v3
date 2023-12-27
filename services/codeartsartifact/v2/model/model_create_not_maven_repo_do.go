package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNotMavenRepoDo struct {

	// 仓库格式
	Format string `json:"format"`

	// 仓库类型
	Type string `json:"type"`

	// 仓库名称
	RepositoryName string `json:"repository_name"`

	// 仓库描述
	Description *string `json:"description,omitempty"`

	// 路径白名单
	IncludesPattern string `json:"includes_pattern"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 共享策略
	ShareRight *string `json:"share_right,omitempty"`
}

func (o CreateNotMavenRepoDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotMavenRepoDo struct{}"
	}

	return strings.Join([]string{"CreateNotMavenRepoDo", string(data)}, " ")
}
