package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKnowledgeSkillResponse Response Object
type UpdateKnowledgeSkillResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateKnowledgeSkillResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKnowledgeSkillResponse struct{}"
	}

	return strings.Join([]string{"UpdateKnowledgeSkillResponse", string(data)}, " ")
}
