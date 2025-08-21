package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainTemplateRequest Request Object
type CreateDomainTemplateRequest struct {
	Body *CreateTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateDomainTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainTemplateRequest", string(data)}, " ")
}
