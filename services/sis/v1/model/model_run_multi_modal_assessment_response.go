package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunMultiModalAssessmentResponse struct {
	// 综合评分，0-100

	Score *float32 `json:"score,omitempty"`
	// 完整性评分，0-100 表示有多少比例的单词发音是清楚的

	Completeness *float32 `json:"completeness,omitempty"`
	// 音频/视频时长，单位秒

	Duration *float32 `json:"duration,omitempty"`

	Pronunciation *Pronunciation `json:"pronunciation,omitempty"`

	Fluency *Fluency `json:"fluency,omitempty"`
	// 单词评测打分表

	Words *[]Word `json:"words,omitempty"`
	// 评测失败时定位问题使用的字段

	TraceId        *string `json:"traceId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunMultiModalAssessmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunMultiModalAssessmentResponse struct{}"
	}

	return strings.Join([]string{"RunMultiModalAssessmentResponse", string(data)}, " ")
}
