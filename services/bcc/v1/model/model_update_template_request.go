package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateRequest Request Object
type UpdateTemplateRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 模板ID
	TemplateId string `json:"template_id"`

	Body *UpdateTemplateBody `json:"body,omitempty"`
}

func (o UpdateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateRequest", string(data)}, " ")
}
