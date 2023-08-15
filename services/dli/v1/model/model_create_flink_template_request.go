package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkTemplateRequest Request Object
type CreateFlinkTemplateRequest struct {
	Body *CreateFlinkTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateFlinkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateFlinkTemplateRequest", string(data)}, " ")
}
