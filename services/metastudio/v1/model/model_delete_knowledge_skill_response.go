package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKnowledgeSkillResponse Response Object
type DeleteKnowledgeSkillResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteKnowledgeSkillResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKnowledgeSkillResponse struct{}"
	}

	return strings.Join([]string{"DeleteKnowledgeSkillResponse", string(data)}, " ")
}
