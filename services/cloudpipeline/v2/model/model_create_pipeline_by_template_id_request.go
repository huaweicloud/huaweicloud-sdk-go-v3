package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePipelineByTemplateIdRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 模板ID
	TemplateId string `json:"template_id"`

	// 语言类型 中文:zh-cn 英文:en-us，默认en-us
	XLanguage *string `json:"X-Language,omitempty"`

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
