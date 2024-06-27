package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestCaseAttachmentInfo 对外附件
type TestCaseAttachmentInfo struct {

	// 附件是否要被覆盖
	Override *bool `json:"override,omitempty"`

	// 文档id
	DocId *string `json:"doc_id,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 文件大小
	FileSize *string `json:"file_size,omitempty"`

	// 重复用例ID
	OverrideId *string `json:"override_id,omitempty"`

	// 相关类型
	RelatedType *string `json:"related_type,omitempty"`

	// 保存文件名
	StoreFileName *string `json:"store_file_name,omitempty"`

	// 系统区分
	SystemType *string `json:"system_type,omitempty"`

	// 区分文件存储系统
	StorageSystem *string `json:"storage_system,omitempty"`
}

func (o TestCaseAttachmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseAttachmentInfo struct{}"
	}

	return strings.Join([]string{"TestCaseAttachmentInfo", string(data)}, " ")
}
