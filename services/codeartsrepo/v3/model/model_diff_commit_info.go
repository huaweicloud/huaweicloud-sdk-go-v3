package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffCommitInfo struct {

	// 变更前文件路径
	OldPath *string `json:"old_path,omitempty"`

	// 变更后文件路径
	NewPath *string `json:"new_path,omitempty"`

	// 变更前文件模式
	AMode *string `json:"a_mode,omitempty"`

	// 变更后文件模式
	BMode *string `json:"b_mode,omitempty"`

	// 此次变更是否新增文件
	NewFile *bool `json:"new_file,omitempty"`

	// 此次变更是否重命名文件
	RenamedFile *bool `json:"renamed_file,omitempty"`

	// 此次变更是否删除文件
	DeletedFile *bool `json:"deleted_file,omitempty"`

	// 差异信息
	Diff *string `json:"diff,omitempty"`
}

func (o DiffCommitInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffCommitInfo struct{}"
	}

	return strings.Join([]string{"DiffCommitInfo", string(data)}, " ")
}
