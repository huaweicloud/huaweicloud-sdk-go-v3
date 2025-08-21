package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRefCompareResponse Response Object
type ShowRefCompareResponse struct {
	Commit *CommitDto `json:"commit,omitempty"`

	// commit相关信息列表
	Commits *[]CommitDto `json:"commits,omitempty"`

	// 变更文件信息
	Diffs *[]DiffDto `json:"diffs,omitempty"`

	// 变更文件信息
	DiffsFiles *[]DiffDto `json:"diffs_files,omitempty"`

	// 是否超时
	CompareTimeout *bool `json:"compare_timeout,omitempty"`

	// 是否相同
	CompareSameRef *bool `json:"compare_same_ref,omitempty"`

	// 冲突文件
	ConflictFiles *[]ConflictFileDto `json:"conflict_files,omitempty"`

	// 增加行数
	AddedLines *int32 `json:"added_lines,omitempty"`

	// 删除行数
	RemovedLines *int32 `json:"removed_lines,omitempty"`

	// 提交数量
	CommitsCount *int32 `json:"commits_count,omitempty"`

	// 文件变更数量
	DiffsCount *int32 `json:"diffs_count,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRefCompareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRefCompareResponse struct{}"
	}

	return strings.Join([]string{"ShowRefCompareResponse", string(data)}, " ")
}
