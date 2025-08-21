package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileChangeDto 文件变更详情
type FileChangeDto struct {

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 变更文件源分支路径
	OldPath *string `json:"old_path,omitempty"`

	// 变更文件目标分支路径
	NewPath *string `json:"new_path,omitempty"`

	// 旧文件权限
	AMode *string `json:"a_mode,omitempty"`

	// 新文件权限
	BMode *string `json:"b_mode,omitempty"`

	// 文件权限是否变更
	ModeChanged *bool `json:"mode_changed,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 是否为新增文件
	NewFile *bool `json:"new_file,omitempty"`

	// 是否为重命名文件
	RenamedFile *bool `json:"renamed_file,omitempty"`

	// 是否为删除文件
	DeletedFile *bool `json:"deleted_file,omitempty"`

	// 变更文件内容
	Diff *string `json:"diff,omitempty"`

	// 是否为二进制文件
	Binary *bool `json:"binary,omitempty"`

	// 是否为大文件
	TooLarge *bool `json:"too_large,omitempty"`

	// 是否折叠文件
	Collapsed *bool `json:"collapsed,omitempty"`

	// 文件新增行
	AddedLines *int32 `json:"added_lines,omitempty"`

	// 文件删除行
	RemovedLines *int32 `json:"removed_lines,omitempty"`

	// 文件最新commitId
	ContentSha *string `json:"content_sha,omitempty"`

	DiffRefs *DiffRefDto `json:"diff_refs,omitempty"`
}

func (o FileChangeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileChangeDto struct{}"
	}

	return strings.Join([]string{"FileChangeDto", string(data)}, " ")
}
