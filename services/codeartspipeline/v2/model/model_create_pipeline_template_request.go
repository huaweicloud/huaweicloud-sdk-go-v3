package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineTemplateRequest Request Object
type CreatePipelineTemplateRequest struct {

	// 租户ID
	TenantId string `json:"tenant_id"`

	Body *PipelineTemplateDto `json:"body,omitempty"`
}

func (o CreatePipelineTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreatePipelineTemplateRequest", string(data)}, " ")
}
