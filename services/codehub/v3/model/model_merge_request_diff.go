package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestDiff struct {

	// base提交
	BaseCommitSha *string `json:"baseCommitSha,omitempty"`

	// 提交数
	CommitsCount *float64 `json:"commitsCount,omitempty"`

	// 创建时间
	CreatedAt *string `json:"createdAt,omitempty"`

	// head提交
	HeadCommitSha *string `json:"headCommitSha,omitempty"`

	// 合并请求id
	MergeRequestId *float64 `json:"mergeRequestId,omitempty"`

	// start提交
	StartCommitSha *string `json:"startCommitSha,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (o MergeRequestDiff) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestDiff struct{}"
	}

	return strings.Join([]string{"MergeRequestDiff", string(data)}, " ")
}
