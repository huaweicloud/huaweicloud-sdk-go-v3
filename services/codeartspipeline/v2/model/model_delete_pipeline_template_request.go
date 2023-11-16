package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipelineTemplateRequest Request Object
type DeletePipelineTemplateRequest struct {

	// 租户ID
	TenantId string `json:"tenant_id"`

	// 模板ID
	TemplateId string `json:"template_id"`
}

func (o DeletePipelineTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeletePipelineTemplateRequest", string(data)}, " ")
}
