package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeLibraryResponse Response Object
type CreateKnowledgeLibraryResponse struct {

	// 知识库ID。
	KnowledgeLibraryId *string `json:"knowledge_library_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKnowledgeLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeLibraryResponse struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeLibraryResponse", string(data)}, " ")
}
