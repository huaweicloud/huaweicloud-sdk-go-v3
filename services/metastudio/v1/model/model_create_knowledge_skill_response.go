package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeSkillResponse Response Object
type CreateKnowledgeSkillResponse struct {

	// 技能ID。
	SkillId *string `json:"skill_id,omitempty"`

	// 技能标识。
	Identify *string `json:"identify,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKnowledgeSkillResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeSkillResponse struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeSkillResponse", string(data)}, " ")
}
