package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEmailTemplateRequest Request Object
type UpdateEmailTemplateRequest struct {
	Body *UpdateEmailTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateEmailTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEmailTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateEmailTemplateRequest", string(data)}, " ")
}
