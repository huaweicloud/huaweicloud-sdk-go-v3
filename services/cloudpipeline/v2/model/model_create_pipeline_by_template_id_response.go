package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineByTemplateIdResponse Response Object
type CreatePipelineByTemplateIdResponse struct {

	// 流水线ID
	PipelineId     *string `json:"pipeline_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePipelineByTemplateIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineByTemplateIdResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineByTemplateIdResponse", string(data)}, " ")
}
