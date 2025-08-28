package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KnowledgeLibraryInfo 知识库基本信息。
type KnowledgeLibraryInfo struct {

	// 知识库ID。
	KnowledgeLibraryId *string `json:"knowledge_library_id,omitempty"`

	// 知识库名称。
	Name *string `json:"name,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	KnowledgeType *KnowledgeTypeEnum `json:"knowledge_type,omitempty"`

	// 知识库大小(文档库为文档数量)
	KnowledgeSize *int32 `json:"knowledge_size,omitempty"`

	// 文档库召回topk
	Topk *int32 `json:"topk,omitempty"`

	// 知识库召回得分
	Score *float64 `json:"score,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o KnowledgeLibraryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KnowledgeLibraryInfo struct{}"
	}

	return strings.Join([]string{"KnowledgeLibraryInfo", string(data)}, " ")
}
