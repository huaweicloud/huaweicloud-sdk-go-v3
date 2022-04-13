package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnalysisModelResponse struct {
	// 分析任务名称，正则：\"^[a-zA-Z][a-zA-Z0-9_]{0,63}$\"

	Name string `json:"name"`
	// 分析任务显示名称，正则：\"^[\\\\u4E00-\\\\u9FA5A-Za-z0-9_@#.-]{1,64}$\"

	DisplayName *string `json:"display_name,omitempty"`
	// 分析任务类型，转换计算（transform）、聚合计算（aggregate）、流计算（stream）

	Type string `json:"type"`

	Transform *TransformModel `json:"transform,omitempty"`

	Aggregate *AggregateModel `json:"aggregate,omitempty"`

	Stream *StreamModel `json:"stream,omitempty"`
	// 分析任务ID

	AnalysisId *string `json:"analysis_id,omitempty"`
}

func (o AnalysisModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisModelResponse struct{}"
	}

	return strings.Join([]string{"AnalysisModelResponse", string(data)}, " ")
}
