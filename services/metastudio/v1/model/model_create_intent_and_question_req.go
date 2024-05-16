package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIntentAndQuestionReq 创建知识库意图和问法请求。
type CreateIntentAndQuestionReq struct {

	// 主题。
	Name string `json:"name"`

	// 问题答案。
	Answer string `json:"answer"`

	// 技能ID。
	SkillId string `json:"skill_id"`

	// 问法列表
	QuestionList *[]KnowledgeQuestionCreateInfo `json:"question_list,omitempty"`
}

func (o CreateIntentAndQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIntentAndQuestionReq struct{}"
	}

	return strings.Join([]string{"CreateIntentAndQuestionReq", string(data)}, " ")
}
