package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSearchAnalysisResponse Response Object
type CreateSearchAnalysisResponse struct {

	// 统计分析结果数据
	Datarows *[][]interface{} `json:"datarows,omitempty"`

	// 统计分析结果字段类型
	Schema *[]AnalysisField `json:"schema,omitempty"`

	// 返回的统计分析结果条数
	Size *int32 `json:"size,omitempty"`

	// 统计分析结果总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateSearchAnalysisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchAnalysisResponse struct{}"
	}

	return strings.Join([]string{"CreateSearchAnalysisResponse", string(data)}, " ")
}
