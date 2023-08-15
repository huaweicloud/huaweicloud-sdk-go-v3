package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopPipelineJobResponse Response Object
type StopPipelineJobResponse struct {

	// 管道ID
	PipelineId     *string `json:"pipeline_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopPipelineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPipelineJobResponse struct{}"
	}

	return strings.Join([]string{"StopPipelineJobResponse", string(data)}, " ")
}
