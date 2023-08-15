package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineTemplateDetailRequest Request Object
type ShowPipelineTemplateDetailRequest struct {

	// 租户ID
	TenantId string `json:"tenant_id"`

	// 模板ID
	TemplateId string `json:"template_id"`
}

func (o ShowPipelineTemplateDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineTemplateDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowPipelineTemplateDetailRequest", string(data)}, " ")
}
