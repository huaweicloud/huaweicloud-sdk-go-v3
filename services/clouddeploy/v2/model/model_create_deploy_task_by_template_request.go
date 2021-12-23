package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeployTaskByTemplateRequest struct {
	Body *TemplateTaskRequestBody `json:"body,omitempty"`
}

func (o CreateDeployTaskByTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeployTaskByTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateDeployTaskByTemplateRequest", string(data)}, " ")
}
