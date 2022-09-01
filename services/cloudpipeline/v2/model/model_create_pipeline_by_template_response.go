package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePipelineByTemplateResponse struct {

	// 实例ID
	TaskId         *string `json:"task_id,omitempty" xml:"task_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePipelineByTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineByTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineByTemplateResponse", string(data)}, " ")
}
