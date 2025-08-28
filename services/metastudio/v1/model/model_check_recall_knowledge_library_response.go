package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRecallKnowledgeLibraryResponse Response Object
type CheckRecallKnowledgeLibraryResponse struct {
	Body *[]RecallKnowledgeLibraryInfo `json:"body,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckRecallKnowledgeLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRecallKnowledgeLibraryResponse struct{}"
	}

	return strings.Join([]string{"CheckRecallKnowledgeLibraryResponse", string(data)}, " ")
}
