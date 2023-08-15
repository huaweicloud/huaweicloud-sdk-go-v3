package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTranscodeTemplateRequest Request Object
type CreateTranscodeTemplateRequest struct {
	Body *CreateTranscodeTemplate `json:"body,omitempty"`
}

func (o CreateTranscodeTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTranscodeTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateTranscodeTemplateRequest", string(data)}, " ")
}
