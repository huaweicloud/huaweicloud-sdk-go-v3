package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeIntentReq 创建知识库意图请求。
type CreateKnowledgeIntentReq struct {

	// 主题。
	Name string `json:"name"`

	// 问题答案。
	Answer string `json:"answer"`

	// 技能ID。
	SkillId string `json:"skill_id"`
}

func (o CreateKnowledgeIntentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeIntentReq struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeIntentReq", string(data)}, " ")
}
