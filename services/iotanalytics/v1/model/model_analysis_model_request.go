package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AnalysisModelRequest 分析任务
type AnalysisModelRequest struct {

	// 分析任务名称，正则：\"^[a-zA-Z][a-zA-Z0-9_]{0,63}$\"
	Name string `json:"name"`

	// 分析任务显示名称，正则：\"^[\\\\u4E00-\\\\u9FA5A-Za-z0-9_@#.-]{1,64}$\"
	DisplayName *string `json:"display_name,omitempty"`

	// 分析任务类型，转换计算（transform）、聚合计算（aggregate）、流计算（stream）
	Type string `json:"type"`

	Transform *TransformModel `json:"transform,omitempty"`

	Aggregate *AggregateModel `json:"aggregate,omitempty"`

	Stream *StreamModel `json:"stream,omitempty"`
}

func (o AnalysisModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisModelRequest struct{}"
	}

	return strings.Join([]string{"AnalysisModelRequest", string(data)}, " ")
}
