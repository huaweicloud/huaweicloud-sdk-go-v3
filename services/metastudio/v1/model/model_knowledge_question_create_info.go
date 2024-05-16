package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KnowledgeQuestionCreateInfo 创建知识库问法请求。
type KnowledgeQuestionCreateInfo struct {

	// 问法。
	Question string `json:"question"`
}

func (o KnowledgeQuestionCreateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KnowledgeQuestionCreateInfo struct{}"
	}

	return strings.Join([]string{"KnowledgeQuestionCreateInfo", string(data)}, " ")
}
