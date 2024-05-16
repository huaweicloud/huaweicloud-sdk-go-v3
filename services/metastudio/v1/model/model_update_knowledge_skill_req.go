package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKnowledgeSkillReq 修改知识库技能请求。
type UpdateKnowledgeSkillReq struct {

	// 技能名称。
	Name *string `json:"name,omitempty"`

	// 技能标识。
	Identify *string `json:"identify,omitempty"`
}

func (o UpdateKnowledgeSkillReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKnowledgeSkillReq struct{}"
	}

	return strings.Join([]string{"UpdateKnowledgeSkillReq", string(data)}, " ")
}
