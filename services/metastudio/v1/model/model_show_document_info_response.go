package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDocumentInfoResponse Response Object
type ShowDocumentInfoResponse struct {

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

	// 分段策略，多个策略之间用逗号分割。 > title:标题分割；separator:分隔符分割
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

func (o ShowDocumentInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDocumentInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowDocumentInfoResponse", string(data)}, " ")
}
