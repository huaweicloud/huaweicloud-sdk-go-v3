package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPipelineJobRequest struct {
	// 管道ID

	PipelineId string `json:"pipeline_id"`
}

func (o ShowPipelineJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineJobRequest struct{}"
	}

	return strings.Join([]string{"ShowPipelineJobRequest", string(data)}, " ")
}
