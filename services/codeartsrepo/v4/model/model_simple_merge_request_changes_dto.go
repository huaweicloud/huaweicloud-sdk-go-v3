package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleMergeRequestChangesDto 合并请求文件变更列表详情
type SimpleMergeRequestChangesDto struct {

	// MR id
	Id *int32 `json:"id,omitempty"`

	// 仓库id
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// 合并请求标题
	Title *string `json:"title,omitempty"`

	// 合并请求描述
	Description *string `json:"description,omitempty"`

	// 合并请求状态
	State *string `json:"state,omitempty"`

	// 合并请求源分支名
	SourceBranch *string `json:"source_branch,omitempty"`

	// 合并请求目标分支名
	TargetBranch *string `json:"target_branch,omitempty"`

	Author *UserBasicDto `json:"author,omitempty"`

	// 合并请求URL链接
	WebUrl *string `json:"web_url,omitempty"`

	// 合并请求创建时间
	CreateAt *string `json:"create_at,omitempty"`

	// 合并请求最后更新时间
	UpdateAt *string `json:"update_at,omitempty"`
}

func (o SimpleMergeRequestChangesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleMergeRequestChangesDto struct{}"
	}

	return strings.Join([]string{"SimpleMergeRequestChangesDto", string(data)}, " ")
}
