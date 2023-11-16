package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipelineTemplateResponse Response Object
type DeletePipelineTemplateResponse struct {
	TemplateId     *string `json:"templateId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePipelineTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeletePipelineTemplateResponse", string(data)}, " ")
}
