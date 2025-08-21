package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestDiffExternalDto struct {

	// 合并请求diff id
	Id *int32 `json:"id,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 合并请求id
	MergeRequestId *int32 `json:"merge_request_id,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// base commit节点
	BaseCommitSha *string `json:"base_commit_sha,omitempty"`

	// diff文件数量
	RealSize *string `json:"real_size,omitempty"`

	// head commit节点
	HeadCommitSha *string `json:"head_commit_sha,omitempty"`

	// start commit节点
	StartCommitSha *string `json:"start_commit_sha,omitempty"`

	// commit数量
	CommitsCount *int32 `json:"commits_count,omitempty"`

	// 外部diff文件
	ExternalDiff *string `json:"external_diff,omitempty"`

	// 外部差异存储
	ExternalDiffStore *int32 `json:"external_diff_store,omitempty"`

	// 是否外部存储
	StoredExternally *bool `json:"stored_externally,omitempty"`

	// 新增行数
	AddedLines *int32 `json:"added_lines,omitempty"`

	// 删除行数
	RemovedLines *int32 `json:"removed_lines,omitempty"`
}

func (o MergeRequestDiffExternalDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestDiffExternalDto struct{}"
	}

	return strings.Join([]string{"MergeRequestDiffExternalDto", string(data)}, " ")
}
