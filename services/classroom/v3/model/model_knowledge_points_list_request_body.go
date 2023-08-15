package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KnowledgePointsListRequestBody 知识点列表请求参数
type KnowledgePointsListRequestBody struct {

	// 名称模糊查询
	Name *string `json:"name,omitempty"`

	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`

	// 起始页
	StartIndex *int32 `json:"start_index,omitempty"`
}

func (o KnowledgePointsListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KnowledgePointsListRequestBody struct{}"
	}

	return strings.Join([]string{"KnowledgePointsListRequestBody", string(data)}, " ")
}
