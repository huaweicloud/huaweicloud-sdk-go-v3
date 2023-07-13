package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMyKnowledgePointsRequest Request Object
type ListMyKnowledgePointsRequest struct {
	Body *KnowledgePointsListRequestBody `json:"body,omitempty"`
}

func (o ListMyKnowledgePointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMyKnowledgePointsRequest struct{}"
	}

	return strings.Join([]string{"ListMyKnowledgePointsRequest", string(data)}, " ")
}
