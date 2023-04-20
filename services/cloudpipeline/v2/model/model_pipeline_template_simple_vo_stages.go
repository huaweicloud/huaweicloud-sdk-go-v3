package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTemplateSimpleVoStages struct {

	// 阶段名称
	Name *string `json:"name,omitempty"`

	// 序列号
	Sequence *int32 `json:"sequence,omitempty"`
}

func (o PipelineTemplateSimpleVoStages) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTemplateSimpleVoStages struct{}"
	}

	return strings.Join([]string{"PipelineTemplateSimpleVoStages", string(data)}, " ")
}
