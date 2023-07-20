package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobTemplatesRequest Request Object
type UpdateJobTemplatesRequest struct {
	TemplateId string `json:"template_id"`

	Body *UpdateJobTemplatesRequestBody `json:"body,omitempty"`
}

func (o UpdateJobTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobTemplatesRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobTemplatesRequest", string(data)}, " ")
}
