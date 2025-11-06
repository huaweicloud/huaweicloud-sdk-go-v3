package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MrConflictFileDto 冲突文件详情
type MrConflictFileDto struct {

	// **参数解释：** 旧文件名称。
	OldPath *string `json:"old_path,omitempty"`

	// **参数解释：** 新文件名称。
	NewPath *string `json:"new_path,omitempty"`

	// blob_icon
	BlobIcon *string `json:"blob_icon,omitempty"`

	// blob 路径
	BlobPath *string `json:"blob_path,omitempty"`

	// 冲突类型
	ConflictType *string `json:"conflict_type,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`

	// 内容路径
	ContentPath *string `json:"content_path,omitempty"`

	// 片段
	Sections *[]ConflictSectionDto `json:"sections,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o MrConflictFileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MrConflictFileDto struct{}"
	}

	return strings.Join([]string{"MrConflictFileDto", string(data)}, " ")
}
