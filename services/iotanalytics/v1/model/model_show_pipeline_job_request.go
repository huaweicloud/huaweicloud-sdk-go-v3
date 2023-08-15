package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineJobRequest Request Object
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
