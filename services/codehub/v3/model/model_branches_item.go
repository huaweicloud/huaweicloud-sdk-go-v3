package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BranchesItem struct {
	Commit *CommitV2 `json:"commit,omitempty"`

	DivergingCommitCounts *DivergingCommitCounts `json:"diverging_commit_counts,omitempty"`

	// 分支名
	Name *string `json:"name,omitempty"`
}

func (o BranchesItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BranchesItem struct{}"
	}

	return strings.Join([]string{"BranchesItem", string(data)}, " ")
}
