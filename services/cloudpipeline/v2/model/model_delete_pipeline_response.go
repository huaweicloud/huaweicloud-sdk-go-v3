package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipelineResponse Response Object
type DeletePipelineResponse struct {

	// 流水线ID
	PipelineId     *string `json:"pipeline_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineResponse struct{}"
	}

	return strings.Join([]string{"DeletePipelineResponse", string(data)}, " ")
}
