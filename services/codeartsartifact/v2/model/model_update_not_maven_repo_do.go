package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNotMavenRepoDo struct {

	// 仓库名称
	RepoName string `json:"repo_name"`

	// 仓库格式
	Format string `json:"format"`

	// 仓库描述
	Description *string `json:"description,omitempty"`

	// 仓库id列表
	RepositoryIds []string `json:"repository_ids"`

	// 路径白名单
	IncludesPattern *string `json:"includes_pattern,omitempty"`

	// 仓库属性-覆盖策略
	DeploymentPolicy *string `json:"deployment_policy,omitempty"`

	// 自动清理快照
	AutoCleanSnapshot *bool `json:"auto_clean_snapshot,omitempty"`

	// 快照保存时间长度
	SnapshotAliveDays *string `json:"snapshot_alive_days,omitempty"`

	// 最大不同快照个数
	MaxUniqueSnapshots *string `json:"max_unique_snapshots,omitempty"`

	// 是否允许匿名
	AllowAnonymous *bool `json:"allow_anonymous,omitempty"`
}

func (o UpdateNotMavenRepoDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotMavenRepoDo struct{}"
	}

	return strings.Join([]string{"UpdateNotMavenRepoDo", string(data)}, " ")
}
