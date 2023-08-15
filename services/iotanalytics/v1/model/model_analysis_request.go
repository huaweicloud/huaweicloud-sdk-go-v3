package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AnalysisRequest 分析任务
type AnalysisRequest struct {

	// 分析任务名，必须是模型中已存在的
	Name string `json:"name"`

	Transform *DtTransformRequest `json:"transform,omitempty"`

	Aggregate *DtAggregateRequest `json:"aggregate,omitempty"`

	Stream *DtStreamRequest `json:"stream,omitempty"`
}

func (o AnalysisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisRequest struct{}"
	}

	return strings.Join([]string{"AnalysisRequest", string(data)}, " ")
}
