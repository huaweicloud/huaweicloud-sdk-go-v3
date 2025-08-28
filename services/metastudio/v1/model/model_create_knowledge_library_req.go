package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeLibraryReq 创建知识库请求。
type CreateKnowledgeLibraryReq struct {

	// 知识库名称。
	Name string `json:"name"`

	Language *LanguageEnum `json:"language"`

	KnowledgeType *KnowledgeTypeEnum `json:"knowledge_type"`

	// 知识库召回数
	Topk *int32 `json:"topk,omitempty"`

	// 知识库召回得分
	Score *float64 `json:"score,omitempty"`
}

func (o CreateKnowledgeLibraryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeLibraryReq struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeLibraryReq", string(data)}, " ")
}
