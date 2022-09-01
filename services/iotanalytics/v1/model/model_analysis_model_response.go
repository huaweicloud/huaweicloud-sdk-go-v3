package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnalysisModelResponse struct {

	// 分析任务名称，正则：\"^[a-zA-Z][a-zA-Z0-9_]{0,63}$\"
	Name string `json:"name" xml:"name"`

	// 分析任务显示名称，正则：\"^[\\\\u4E00-\\\\u9FA5A-Za-z0-9_@#.-]{1,64}$\"
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 分析任务类型，转换计算（transform）、聚合计算（aggregate）、流计算（stream）
	Type string `json:"type" xml:"type"`

	Transform *TransformModel `json:"transform,omitempty" xml:"transform"`

	Aggregate *AggregateModel `json:"aggregate,omitempty" xml:"aggregate"`

	Stream *StreamModel `json:"stream,omitempty" xml:"stream"`

	// 分析任务ID
	AnalysisId *string `json:"analysis_id,omitempty" xml:"analysis_id"`
}

func (o AnalysisModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisModelResponse struct{}"
	}

	return strings.Join([]string{"AnalysisModelResponse", string(data)}, " ")
}
