package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecallKnowledgeLibraryInfo 知识库召回测试响应。
type RecallKnowledgeLibraryInfo struct {

	// 文档ID。
	DocumentId *string `json:"document_id,omitempty"`

	// 文档名称。
	FileName *string `json:"file_name,omitempty"`

	// 文档类型。
	FileType *string `json:"file_type,omitempty"`

	// 问答对ID。
	QuestionAnswerId *string `json:"question_answer_id,omitempty"`

	// 召回文本
	Content *string `json:"content,omitempty"`

	// 知识库召回得分
	Score *float64 `json:"score,omitempty"`
}

func (o RecallKnowledgeLibraryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecallKnowledgeLibraryInfo struct{}"
	}

	return strings.Join([]string{"RecallKnowledgeLibraryInfo", string(data)}, " ")
}
