package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeployTaskByTemplateRequest struct {
	Body *TemplateTaskRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateDeployTaskByTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeployTaskByTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateDeployTaskByTemplateRequest", string(data)}, " ")
}
