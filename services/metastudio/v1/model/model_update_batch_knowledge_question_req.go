package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBatchKnowledgeQuestionReq 批量修改知识库问法请求。
type UpdateBatchKnowledgeQuestionReq struct {

	// 问法列表
	QuestionList []KnowledgeQuestionUpdateInfo `json:"question_list"`
}

func (o UpdateBatchKnowledgeQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBatchKnowledgeQuestionReq struct{}"
	}

	return strings.Join([]string{"UpdateBatchKnowledgeQuestionReq", string(data)}, " ")
}
