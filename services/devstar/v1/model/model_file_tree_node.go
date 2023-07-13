package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileTreeNode struct {

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 是否文件夹
	FileType *string `json:"file_type,omitempty"`
}

func (o FileTreeNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileTreeNode struct{}"
	}

	return strings.Join([]string{"FileTreeNode", string(data)}, " ")
}
