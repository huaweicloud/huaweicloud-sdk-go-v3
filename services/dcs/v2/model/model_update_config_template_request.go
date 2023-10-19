package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigTemplateRequest Request Object
type UpdateConfigTemplateRequest struct {

	// 模板ID
	TemplateId string `json:"template_id"`

	Body *UpdateCustomTemplateBody `json:"body,omitempty"`
}

func (o UpdateConfigTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigTemplateRequest", string(data)}, " ")
}
