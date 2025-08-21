package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffDto struct {

	// 旧文件
	OldPath *string `json:"old_path,omitempty"`

	// 新文件
	NewPath *string `json:"new_path,omitempty"`

	// 旧文件类型
	AMode *string `json:"a_mode,omitempty"`

	// 新文件类型
	BMode *string `json:"b_mode,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 是否新增
	NewFile *bool `json:"new_file,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 是否重命名
	RenamedFile *bool `json:"renamed_file,omitempty"`

	// 是否删除文件
	DeletedFile *bool `json:"deleted_file,omitempty"`

	// 比较结果
	Diff *string `json:"diff,omitempty"`

	// 是否二进制
	Binary *bool `json:"binary,omitempty"`

	// 是否过大
	TooLarge *bool `json:"too_large,omitempty"`

	// 是否折叠
	Collapsed *bool `json:"collapsed,omitempty"`

	LineCount *[]int32 `json:"line_count,omitempty"`

	// 增加行数
	AddedLines *int32 `json:"added_lines,omitempty"`

	// 删除行数
	RemovedLines *int32 `json:"removed_lines,omitempty"`
}

func (o DiffDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffDto struct{}"
	}

	return strings.Join([]string{"DiffDto", string(data)}, " ")
}
