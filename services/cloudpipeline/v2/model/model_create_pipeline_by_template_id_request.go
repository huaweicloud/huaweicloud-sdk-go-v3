package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineByTemplateIdRequest Request Object
type CreatePipelineByTemplateIdRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 模板ID
	TemplateId string `json:"template_id"`

	// 微服务ID
	ComponentId *string `json:"component_id,omitempty"`

	Body *PipelineByTemplateDto `json:"body,omitempty"`
}

func (o CreatePipelineByTemplateIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineByTemplateIdRequest struct{}"
	}

	return strings.Join([]string{"CreatePipelineByTemplateIdRequest", string(data)}, " ")
}
