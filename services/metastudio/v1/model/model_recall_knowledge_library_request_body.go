package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecallKnowledgeLibraryRequestBody 知识库召回测试请求。
type RecallKnowledgeLibraryRequestBody struct {

	// 知识库名称。
	KnowledgeLibraryId string `json:"knowledge_library_id"`

	KnowledgeType *KnowledgeTypeEnum `json:"knowledge_type"`

	// 知识库召回请求文本
	Query string `json:"query"`

	// 文档库召回topk
	Topk *int32 `json:"topk,omitempty"`

	// 知识库召回得分
	Score *float64 `json:"score,omitempty"`
}

func (o RecallKnowledgeLibraryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecallKnowledgeLibraryRequestBody struct{}"
	}

	return strings.Join([]string{"RecallKnowledgeLibraryRequestBody", string(data)}, " ")
}
