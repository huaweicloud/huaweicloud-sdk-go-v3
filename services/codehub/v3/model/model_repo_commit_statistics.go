package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 仓库统计信息
type RepoCommitStatistics struct {

	// 仓库总提交次数
	AllBranchCommitsCount *int32 `json:"all_branch_commits_count,omitempty" xml:"all_branch_commits_count"`

	// 近15日每日代码提交行数
	Codelines *[]RepoDailyCodeline `json:"codelines,omitempty" xml:"codelines"`

	// 对应分支仓库总提交次数
	Count *int32 `json:"count,omitempty" xml:"count"`

	Event *RepoStatisticsEvent `json:"event,omitempty" xml:"event"`

	// 仓库统计列表
	Statistics *[]RepoStatistics `json:"statistics,omitempty" xml:"statistics"`

	// 仓库统计次数
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o RepoCommitStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoCommitStatistics struct{}"
	}

	return strings.Join([]string{"RepoCommitStatistics", string(data)}, " ")
}
