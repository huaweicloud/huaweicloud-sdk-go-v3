package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentInfo 附件信息
type AttachmentInfo struct {

	// 文档id
	DocId *string `json:"doc_id,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 保存文件名
	StoreFileName *string `json:"store_file_name,omitempty"`

	// 附件类型 0 本地上传  other 关联文档
	RelatedType *string `json:"related_type,omitempty"`

	// 文件大小
	FileSize *int32 `json:"file_size,omitempty"`
}

func (o AttachmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentInfo struct{}"
	}

	return strings.Join([]string{"AttachmentInfo", string(data)}, " ")
}
