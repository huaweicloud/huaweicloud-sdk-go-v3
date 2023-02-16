package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
