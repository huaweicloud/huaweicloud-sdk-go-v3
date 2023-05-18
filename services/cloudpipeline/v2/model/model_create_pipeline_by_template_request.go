package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePipelineByTemplateRequest struct {
	Body *TemplateCddl `json:"body,omitempty"`
}

func (o CreatePipelineByTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineByTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreatePipelineByTemplateRequest", string(data)}, " ")
}
