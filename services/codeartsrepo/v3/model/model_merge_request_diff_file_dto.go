package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestDiffFileDto struct {

	// 合并请求head sha
	ContentSha *string `json:"content_sha,omitempty"`

	// 是否为submodule
	Submodule *bool `json:"submodule,omitempty"`

	// 内容是否扩展
	Expanded *bool `json:"expanded,omitempty"`

	DiffRefs *DiffRefsDto `json:"diff_refs,omitempty"`

	// mode是否修改
	ModeChanged *bool `json:"mode_changed,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 旧路径
	OldPath *string `json:"old_path,omitempty"`

	// 新路径
	NewPath *string `json:"new_path,omitempty"`

	// 旧mode
	AMode *string `json:"a_mode,omitempty"`

	// 新mode
	BMode *string `json:"b_mode,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 是否为新增文件
	NewFile *bool `json:"new_file,omitempty"`

	// 是否为重命名文件
	RenamedFile *bool `json:"renamed_file,omitempty"`

	// 是否为删除文件
	DeletedFile *bool `json:"deleted_file,omitempty"`

	// 文件变更内容
	Diff *string `json:"diff,omitempty"`

	// 是否为二进制文件
	Binary *bool `json:"binary,omitempty"`

	// 是否过大
	TooLarge *bool `json:"too_large,omitempty"`

	// 是否折叠
	Collapsed *bool `json:"collapsed,omitempty"`

	// 单个文件可取行数范围
	LineCount *[]int32 `json:"line_count,omitempty"`

	// 新增行数
	AddedLines *int32 `json:"added_lines,omitempty"`

	// 删除行数
	RemovedLines *int32 `json:"removed_lines,omitempty"`

	// 文件blob_id
	BlobId *string `json:"blob_id,omitempty"`
}

func (o MergeRequestDiffFileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestDiffFileDto struct{}"
	}

	return strings.Join([]string{"MergeRequestDiffFileDto", string(data)}, " ")
}
