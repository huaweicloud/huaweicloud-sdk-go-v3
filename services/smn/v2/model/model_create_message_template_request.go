package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMessageTemplateRequest struct {
	Body *CreateMessageTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateMessageTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateMessageTemplateRequest", string(data)}, " ")
}
