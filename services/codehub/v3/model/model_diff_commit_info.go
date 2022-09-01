package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffCommitInfo struct {

	// 变更前文件路径
	OldPath *string `json:"old_path,omitempty" xml:"old_path"`

	// 变更后文件路径
	NewPath *string `json:"new_path,omitempty" xml:"new_path"`

	// 变更前文件模式
	AMode *string `json:"a_mode,omitempty" xml:"a_mode"`

	// 变更后文件模式
	BMode *string `json:"b_mode,omitempty" xml:"b_mode"`

	// 此次变更是否新增文件
	NewFile *bool `json:"new_file,omitempty" xml:"new_file"`

	// 此次变更是否重命名文件
	RenamedFile *bool `json:"renamed_file,omitempty" xml:"renamed_file"`

	// 此次变更是否删除文件
	DeletedFile *bool `json:"deleted_file,omitempty" xml:"deleted_file"`

	// 差异信息
	Diff *bool `json:"diff,omitempty" xml:"diff"`
}

func (o DiffCommitInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffCommitInfo struct{}"
	}

	return strings.Join([]string{"DiffCommitInfo", string(data)}, " ")
}
