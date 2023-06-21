package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMyKnowledgePointsResponse struct {

	// 知识点数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 知识点信息
	Data           *[]KnowledgePointInfo `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListMyKnowledgePointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMyKnowledgePointsResponse struct{}"
	}

	return strings.Join([]string{"ListMyKnowledgePointsResponse", string(data)}, " ")
}
