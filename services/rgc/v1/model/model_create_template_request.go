package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateRequest Request Object
type CreateTemplateRequest struct {
	Body *CreateTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateRequest", string(data)}, " ")
}
