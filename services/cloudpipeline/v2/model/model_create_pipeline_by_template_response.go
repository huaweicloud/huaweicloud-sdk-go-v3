package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineByTemplateResponse Response Object
type CreatePipelineByTemplateResponse struct {

	// 实例ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePipelineByTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineByTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineByTemplateResponse", string(data)}, " ")
}
