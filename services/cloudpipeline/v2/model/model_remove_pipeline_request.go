package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RemovePipelineRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us，默认en-us

	XLanguage *string `json:"X-Language,omitempty"`
	// 要删除的流水线ID

	PipelineId string `json:"pipeline_id"`
}

func (o RemovePipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemovePipelineRequest struct{}"
	}

	return strings.Join([]string{"RemovePipelineRequest", string(data)}, " ")
}
