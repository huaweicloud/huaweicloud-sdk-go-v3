package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableAnalysisResponse Response Object
type CreateTableAnalysisResponse struct {

	// 查询结果
	Schema *[]SearchQueryField `json:"schema,omitempty"`

	// 查询结果行
	Datarows *[][]interface{} `json:"datarows,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 总数
	Size *int32 `json:"size,omitempty"`

	// 结果列表
	Results        *[]SearchQueryResult `json:"results,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateTableAnalysisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableAnalysisResponse struct{}"
	}

	return strings.Join([]string{"CreateTableAnalysisResponse", string(data)}, " ")
}
