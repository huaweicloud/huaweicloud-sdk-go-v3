package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分析任务
type AnalysisRequest struct {

	// 分析任务名，必须是模型中已存在的
	Name string `json:"name" xml:"name"`

	Transform *DtTransformRequest `json:"transform,omitempty" xml:"transform"`

	Aggregate *DtAggregateRequest `json:"aggregate,omitempty" xml:"aggregate"`

	Stream *DtStreamRequest `json:"stream,omitempty" xml:"stream"`
}

func (o AnalysisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisRequest struct{}"
	}

	return strings.Join([]string{"AnalysisRequest", string(data)}, " ")
}
