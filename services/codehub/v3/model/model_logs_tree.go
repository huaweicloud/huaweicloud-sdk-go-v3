package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogsTree struct {

	// 存储块id
	BlobId *string `json:"blob_id,omitempty" xml:"blob_id"`

	Commit *Commit `json:"commit,omitempty" xml:"commit"`

	// 文件名称
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty" xml:"file_path"`

	// MD5
	Md5 *string `json:"md5,omitempty" xml:"md5"`

	// 存储类型
	Type *string `json:"type,omitempty" xml:"type"`
}

func (o LogsTree) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogsTree struct{}"
	}

	return strings.Join([]string{"LogsTree", string(data)}, " ")
}
