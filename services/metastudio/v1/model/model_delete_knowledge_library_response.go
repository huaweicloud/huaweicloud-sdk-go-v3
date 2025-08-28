package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKnowledgeLibraryResponse Response Object
type DeleteKnowledgeLibraryResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteKnowledgeLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKnowledgeLibraryResponse struct{}"
	}

	return strings.Join([]string{"DeleteKnowledgeLibraryResponse", string(data)}, " ")
}
