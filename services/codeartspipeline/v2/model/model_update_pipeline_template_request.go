package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipelineTemplateRequest Request Object
type UpdatePipelineTemplateRequest struct {

	// 租户ID
	TenantId string `json:"tenant_id"`

	// 模板ID
	TemplateId string `json:"template_id"`

	Body *PipelineTemplateDto `json:"body,omitempty"`
}

func (o UpdatePipelineTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipelineTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdatePipelineTemplateRequest", string(data)}, " ")
}
