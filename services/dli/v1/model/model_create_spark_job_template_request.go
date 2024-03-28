package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSparkJobTemplateRequest Request Object
type CreateSparkJobTemplateRequest struct {
	Body *CreateSparkJobTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateSparkJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateSparkJobTemplateRequest", string(data)}, " ")
}
