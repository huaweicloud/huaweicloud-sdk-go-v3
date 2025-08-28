package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKnowledgeLibraryResponse Response Object
type UpdateKnowledgeLibraryResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateKnowledgeLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKnowledgeLibraryResponse struct{}"
	}

	return strings.Join([]string{"UpdateKnowledgeLibraryResponse", string(data)}, " ")
}
