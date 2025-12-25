package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AnalysisResults 分析结果
type AnalysisResults struct {

	// 统计分析结果数据
	Datarows [][]interface{} `json:"datarows"`

	// 统计分析结果字段类型
	Schema []AnalysisField `json:"schema"`

	// 返回的统计分析结果条数
	Size int32 `json:"size"`

	// 统计分析结果总数
	Total int32 `json:"total"`
}

func (o AnalysisResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisResults struct{}"
	}

	return strings.Join([]string{"AnalysisResults", string(data)}, " ")
}
