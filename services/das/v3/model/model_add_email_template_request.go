package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEmailTemplateRequest Request Object
type AddEmailTemplateRequest struct {
	Body *AddEmailTemplateRequestBody `json:"body,omitempty"`
}

func (o AddEmailTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEmailTemplateRequest struct{}"
	}

	return strings.Join([]string{"AddEmailTemplateRequest", string(data)}, " ")
}
