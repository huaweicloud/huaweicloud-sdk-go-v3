package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyTemplateRequest Request Object
type CreatePolicyTemplateRequest struct {
	Body *CreatePolicyTemplateReq `json:"body,omitempty"`
}

func (o CreatePolicyTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreatePolicyTemplateRequest", string(data)}, " ")
}
