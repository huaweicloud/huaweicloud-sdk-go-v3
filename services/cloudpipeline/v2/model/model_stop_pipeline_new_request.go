package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopPipelineNewRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us，默认en-us
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 流水线ID
	PipelineId string `json:"pipeline_id" xml:"pipeline_id"`

	// 流水线执行ID
	BuildId string `json:"build_id" xml:"build_id"`
}

func (o StopPipelineNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPipelineNewRequest struct{}"
	}

	return strings.Join([]string{"StopPipelineNewRequest", string(data)}, " ")
}
