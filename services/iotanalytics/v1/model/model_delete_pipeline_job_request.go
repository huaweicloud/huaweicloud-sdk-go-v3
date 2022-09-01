package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePipelineJobRequest struct {

	// 管道ID
	PipelineId string `json:"pipeline_id" xml:"pipeline_id"`
}

func (o DeletePipelineJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineJobRequest struct{}"
	}

	return strings.Join([]string{"DeletePipelineJobRequest", string(data)}, " ")
}
