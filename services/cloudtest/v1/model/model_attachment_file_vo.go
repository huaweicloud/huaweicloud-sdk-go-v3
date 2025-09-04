package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentFileVo 实际的数据类型：单个对象，集合 或 NULL
type AttachmentFileVo struct {

	// 附件Uri
	Uri *string `json:"uri,omitempty"`

	// 文件路径
	Path *string `json:"path,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 文档名
	DocName *string `json:"doc_name,omitempty"`

	// 保存文件名
	StoreName *string `json:"store_name,omitempty"`

	// 文档id
	DocId *int32 `json:"doc_id,omitempty"`

	// 文档大小
	DocSize *string `json:"doc_size,omitempty"`

	// 相关类型
	RelatedType *string `json:"related_type,omitempty"`
}

func (o AttachmentFileVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentFileVo struct{}"
	}

	return strings.Join([]string{"AttachmentFileVo", string(data)}, " ")
}
