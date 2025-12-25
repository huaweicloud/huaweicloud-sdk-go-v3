package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchKnowledgeBaseRequest Request Object
type SearchKnowledgeBaseRequest struct {
	Body *KnowledgeBaseRetrievalReq `json:"body,omitempty"`
}

func (o SearchKnowledgeBaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchKnowledgeBaseRequest struct{}"
	}

	return strings.Join([]string{"SearchKnowledgeBaseRequest", string(data)}, " ")
}
