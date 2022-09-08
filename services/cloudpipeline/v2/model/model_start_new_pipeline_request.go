package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartNewPipelineRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us，默认en-us
	XLanguage *string `json:"X-Language,omitempty"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	Body *StartPipelineParameters `json:"body,omitempty"`
}

func (o StartNewPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartNewPipelineRequest struct{}"
	}

	return strings.Join([]string{"StartNewPipelineRequest", string(data)}, " ")
}
