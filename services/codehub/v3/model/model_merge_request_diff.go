package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestDiff struct {

	// base提交
	BaseCommitSha *string `json:"base_commit_sha,omitempty"`

	// 提交数
	CommitsCount *float64 `json:"commits_count,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// head提交
	HeadCommitSha *string `json:"head_commit_sha,omitempty"`

	// 合并请求id
	MergeRequestId *float64 `json:"merge_request_id,omitempty"`

	// start提交
	StartCommitSha *string `json:"start_commit_sha,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o MergeRequestDiff) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestDiff struct{}"
	}

	return strings.Join([]string{"MergeRequestDiff", string(data)}, " ")
}
