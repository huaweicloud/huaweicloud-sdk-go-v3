package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilesResponseInfo struct {
	// 文件名称

	FileName *string `json:"file_name,omitempty"`
	// 文件路径

	FilePath *string `json:"file_path,omitempty"`
	// 文件大小

	Size *string `json:"size,omitempty"`
	// 编码类型

	Encoding *string `json:"encoding,omitempty"`
	// 分支名称

	Ref *string `json:"ref,omitempty"`
	// 文件块id

	BlobId *string `json:"blob_id,omitempty"`
	// 文件类型

	FileType *string `json:"file_type,omitempty"`
	// 文件内容

	Content *string `json:"content,omitempty"`
}

func (o FilesResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilesResponseInfo struct{}"
	}

	return strings.Join([]string{"FilesResponseInfo", string(data)}, " ")
}
