package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilesResponseInfo struct {

	// 文件名称
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty" xml:"file_path"`

	// 文件大小
	Size *string `json:"size,omitempty" xml:"size"`

	// 编码类型
	Encoding *string `json:"encoding,omitempty" xml:"encoding"`

	// 分支名称
	Ref *string `json:"ref,omitempty" xml:"ref"`

	// 文件块id
	BlobId *string `json:"blob_id,omitempty" xml:"blob_id"`

	// 文件类型
	FileType *string `json:"file_type,omitempty" xml:"file_type"`

	// 文件内容
	Content *string `json:"content,omitempty" xml:"content"`
}

func (o FilesResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilesResponseInfo struct{}"
	}

	return strings.Join([]string{"FilesResponseInfo", string(data)}, " ")
}
