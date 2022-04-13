package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchShowPipelinesStatusRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us，默认en-us

	XLanguage *string `json:"X-Language,omitempty"`
	// 要获取状态的流水线ID，用逗号隔开

	PipelineIds string `json:"pipeline_ids"`
}

func (o BatchShowPipelinesStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowPipelinesStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchShowPipelinesStatusRequest", string(data)}, " ")
}
