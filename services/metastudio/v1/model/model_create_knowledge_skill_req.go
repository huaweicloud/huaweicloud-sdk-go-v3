package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeSkillReq 创建知识库技能请求。
type CreateKnowledgeSkillReq struct {

	// 技能名称。
	Name string `json:"name"`

	// 技能标识。
	Identify string `json:"identify"`
}

func (o CreateKnowledgeSkillReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeSkillReq struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeSkillReq", string(data)}, " ")
}
