package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineTemplateResponse Response Object
type CreatePipelineTemplateResponse struct {
	TemplateId     *string `json:"templateId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePipelineTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineTemplateResponse", string(data)}, " ")
}
