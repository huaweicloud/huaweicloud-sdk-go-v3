package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RepoCommitStatistics 仓库统计信息
type RepoCommitStatistics struct {

	// 仓库总提交次数
	AllBranchCommitsCount *int32 `json:"all_branch_commits_count,omitempty"`

	// 近15日每日代码提交行数
	Codelines *[]RepoDailyCodeline `json:"codelines,omitempty"`

	// 对应分支仓库总提交次数
	Count *int32 `json:"count,omitempty"`

	Event *RepoStatisticsEvent `json:"event,omitempty"`

	// 仓库统计列表
	Statistics *[]RepoStatistics `json:"statistics,omitempty"`

	// 仓库统计次数
	Total *int32 `json:"total,omitempty"`
}

func (o RepoCommitStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoCommitStatistics struct{}"
	}

	return strings.Join([]string{"RepoCommitStatistics", string(data)}, " ")
}
