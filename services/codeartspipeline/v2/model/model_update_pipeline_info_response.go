package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipelineInfoResponse Response Object
type UpdatePipelineInfoResponse struct {

	// 流水线ID
	PipelineId     *string `json:"pipeline_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePipelineInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipelineInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdatePipelineInfoResponse", string(data)}, " ")
}
