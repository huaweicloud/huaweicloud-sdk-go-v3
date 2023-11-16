package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineGroupBindDtoPipelines struct {

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 流水线名
	PipelineName string `json:"pipeline_name"`
}

func (o PipelineGroupBindDtoPipelines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineGroupBindDtoPipelines struct{}"
	}

	return strings.Join([]string{"PipelineGroupBindDtoPipelines", string(data)}, " ")
}
