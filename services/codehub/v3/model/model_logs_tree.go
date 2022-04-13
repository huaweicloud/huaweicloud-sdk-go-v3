package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogsTree struct {
	// 存储块id

	BlobId *string `json:"blob_id,omitempty"`

	Commit *Commit `json:"commit,omitempty"`
	// 文件名称

	FileName *string `json:"file_name,omitempty"`
	// 文件路径

	FilePath *string `json:"file_path,omitempty"`
	// MD5

	Md5 *string `json:"md5,omitempty"`
	// 存储类型

	Type *string `json:"type,omitempty"`
}

func (o LogsTree) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogsTree struct{}"
	}

	return strings.Join([]string{"LogsTree", string(data)}, " ")
}
