package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQuestionAnswerReq 创建问答对请求。
type CreateQuestionAnswerReq struct {

	// 标准问题。
	Question string `json:"question"`

	// 问题答案。
	Answer string `json:"answer"`

	// 知识库ID。
	KnowledgeLibraryId string `json:"knowledge_library_id"`

	// 所有相似问题，多个相似问题间用换行符\\n分隔。
	SimilarQuestions *string `json:"similar_questions,omitempty"`
}

func (o CreateQuestionAnswerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQuestionAnswerReq struct{}"
	}

	return strings.Join([]string{"CreateQuestionAnswerReq", string(data)}, " ")
}
