package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDocumentResponse Response Object
type CreateDocumentResponse struct {

	// 文档ID。
	DocumentId *string `json:"document_id,omitempty"`

	// 知识库ID。
	KnowledgeLibraryId *string `json:"knowledge_library_id,omitempty"`

	// 文档名称。
	FileName *string `json:"file_name,omitempty"`

	// 文档大小，单位字节
	FileSize *int64 `json:"file_size,omitempty"`

	// 文档类型。
	FileType *string `json:"file_type,omitempty"`

	// 分段类型 * 1: 自动分段 * 2: 手动分段
	SplitType *int32 `json:"split_type,omitempty"`

	// 分段长度。
	ChunkSize *int32 `json:"chunk_size,omitempty"`

	// 分段策略，如果添加多个策略用逗号隔开。 取值示例： - title：用标题把一段话分割为多个段落。 - separator：用分隔符把一段话分割为多个段落。
	ChunkType *string `json:"chunk_type,omitempty"`

	// 分隔符
	Separators *[]string `json:"separators,omitempty"`

	// 文档创建时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 文档更新时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	DocumentTaskInfo *DocumentTaskInfo `json:"document_task_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocumentResponse struct{}"
	}

	return strings.Join([]string{"CreateDocumentResponse", string(data)}, " ")
}
