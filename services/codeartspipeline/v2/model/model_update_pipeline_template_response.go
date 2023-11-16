package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipelineTemplateResponse Response Object
type UpdatePipelineTemplateResponse struct {
	TemplateId     *string `json:"templateId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePipelineTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipelineTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdatePipelineTemplateResponse", string(data)}, " ")
}
