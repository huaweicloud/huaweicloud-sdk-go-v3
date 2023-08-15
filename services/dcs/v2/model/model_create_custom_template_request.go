package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomTemplateRequest Request Object
type CreateCustomTemplateRequest struct {
	Body *CreateCustomTemplateBody `json:"body,omitempty"`
}

func (o CreateCustomTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomTemplateRequest", string(data)}, " ")
}
