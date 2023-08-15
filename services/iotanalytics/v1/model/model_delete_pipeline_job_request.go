package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipelineJobRequest Request Object
type DeletePipelineJobRequest struct {

	// 管道ID
	PipelineId string `json:"pipeline_id"`
}

func (o DeletePipelineJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineJobRequest struct{}"
	}

	return strings.Join([]string{"DeletePipelineJobRequest", string(data)}, " ")
}
