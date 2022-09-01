package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnalysisResponse struct {

	// 分析任务ID
	AnalysisId *string `json:"analysis_id,omitempty" xml:"analysis_id"`

	// 分析任务名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 分析任务显示名称
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 分析任务类型，转换计算（transform）、聚合计算（aggregate）、流计算（stream）
	Type *string `json:"type,omitempty" xml:"type"`

	Transform *TransformResponse `json:"transform,omitempty" xml:"transform"`

	Aggregate *AggregateResponse `json:"aggregate,omitempty" xml:"aggregate"`

	Stream *StreamResponse `json:"stream,omitempty" xml:"stream"`
}

func (o AnalysisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisResponse struct{}"
	}

	return strings.Join([]string{"AnalysisResponse", string(data)}, " ")
}
