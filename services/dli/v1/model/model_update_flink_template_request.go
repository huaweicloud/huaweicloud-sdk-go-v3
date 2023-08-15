package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkTemplateRequest Request Object
type UpdateFlinkTemplateRequest struct {
	TemplateId int64 `json:"template_id"`

	Body *UpdateFlinkTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateFlinkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlinkTemplateRequest", string(data)}, " ")
}
