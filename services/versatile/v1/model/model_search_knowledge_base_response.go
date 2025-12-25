package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchKnowledgeBaseResponse Response Object
type SearchKnowledgeBaseResponse struct {

	// **参数解释**： 检索结果总数。  **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 检索结果列表。  **取值范围**： 不涉及。
	RetrieveResultList *[]RetrievalResultInfo `json:"retrieve_result_list,omitempty"`
	HttpStatusCode     int                    `json:"-"`
}

func (o SearchKnowledgeBaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchKnowledgeBaseResponse struct{}"
	}

	return strings.Join([]string{"SearchKnowledgeBaseResponse", string(data)}, " ")
}
